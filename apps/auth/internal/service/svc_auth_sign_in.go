package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/do"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, _ error) {
	resp = &pb_auth.SignInResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}}
	var (
		q         = entity.NewMysqlQuery()
		signIn    *do.SignIn
		server    *pb_auth.ServerInfo
		onOffResp *pb_chat_member.ChatMemberOnOffLineResp
	)
	switch req.AccountType {
	case pb_enum.ACCOUNT_TYPE_MOBILE:
		q.SetFilter("mobile = ?", req.Account)
	case pb_enum.ACCOUNT_TYPE_LARK:
		q.SetFilter("lark_id = ?", req.Account)
	default:
		// 登录类型错误
		resp.Set(ERROR_CODE_AUTH_ACCOUNT_TYPE_ERR, ERROR_AUTH_ACCOUNT_TYPE_ERR)
		return
	}
	// 前端传密码MD5值,服务进行二次MD5加密
	q.SetFilter("password = ?", utils.MD5(req.Password))
	signIn = s.signInTransaction(q, req.Platform)
	if signIn.Err != nil || signIn.Code > 0 {
		resp.Set(signIn.Code, signIn.Msg)
		return
	}
	server = s.getWsServer()
	onOffResp = s.chatMemberOnOffLine(signIn.User.Uid, int64(server.ServerId), req.Platform)
	if onOffResp == nil {
		resp.Set(ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE, ERROR_AUTH_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE, ERROR_AUTH_GRPC_SERVICE_FAILURE)
		return
	}
	copier.Copy(resp.UserInfo, signIn.User)
	copier.Copy(resp.UserInfo.Avatar, signIn.Avatar)
	resp.AccessToken = signIn.AccessToken
	resp.RefreshToken = signIn.RefreshToken
	resp.Server = server
	return
}

func (s *authService) signInTransaction(q *entity.MysqlQuery, platform pb_enum.PLATFORM_TYPE) (signIn *do.SignIn) {
	signIn = new(do.SignIn)
	signIn.User, signIn.Err = s.authRepo.VerifyIdentity(q)
	if signIn.Err != nil {
		signIn.Code = ERROR_CODE_AUTH_QUERY_DB_FAILED
		signIn.Msg = ERROR_AUTH_QUERY_DB_FAILED
		xlog.Warn(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED, signIn.Err.Error())
		return
	}
	if signIn.User.Uid == 0 {
		signIn.Code = ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR
		signIn.Msg = ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR
		return
	}
	q.Reset()
	q.SetFilter("owner_id=?", signIn.User.Uid)
	q.SetFilter("owner_type=?", int32(pb_enum.AVATAR_OWNER_USER_AVATAR))
	signIn.Avatar, signIn.Err = s.avatarRepo.Avatar(q)
	if signIn.Err != nil {
		signIn.Code = ERROR_CODE_AUTH_QUERY_DB_FAILED
		signIn.Msg = ERROR_AUTH_QUERY_DB_FAILED
		xlog.Warn(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED, signIn.Err.Error())
		return
	}
	signIn.AccessToken, signIn.RefreshToken, signIn.Err = s.createToken(signIn.User.Uid, platform)
	if signIn.Err != nil {
		signIn.Code = ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED
		signIn.Msg = ERROR_AUTH_GENERATE_TOKEN_FAILED
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, signIn.Err.Error())
		return
	}
	return
}
