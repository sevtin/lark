package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, _ error) {
	resp = &pb_auth.SignInResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}}
	var (
		w      = entity.NewMysqlWhere()
		user   *po.User
		avatar *po.Avatar
		err    error
	)
	switch req.AccountType {
	case pb_enum.ACCOUNT_TYPE_MOBILE:
		w.SetFilter("mobile = ?", req.Account)
	case pb_enum.ACCOUNT_TYPE_LARK:
		w.SetFilter("lark_id = ?", req.Account)
	default:
		// 登录类型错误
		resp.Set(ERROR_CODE_AUTH_ACCOUNT_TYPE_ERR, ERROR_AUTH_ACCOUNT_TYPE_ERR)
		return
	}
	w.SetFilter("password = ?", utils.MD5(req.Password))
	user, err = s.authRepo.VerifyIdentity(w)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED, err.Error())
		return
	}
	if user.Uid == 0 {
		resp.Set(ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR, ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR)
		return
	}
	w.Reset()
	w.SetFilter("owner_id=?", user.Uid)
	w.SetFilter("owner_type=?", int32(pb_enum.AVATAR_OWNER_USER_AVATAR))
	avatar, err = s.avatarRepo.Avatar(w)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED, err.Error())
		return
	}
	var (
		accessToken  *pb_auth.Token
		refreshToken *pb_auth.Token
	)
	accessToken, refreshToken, err = s.CreateToken(user.Uid, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	copier.Copy(resp.UserInfo.Avatar, avatar)
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	xants.Submit(func() {
		s.userCache.SetUserInfo(resp.UserInfo)
	})
	return
}
