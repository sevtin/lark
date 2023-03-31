package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, _ error) {
	resp = &pb_auth.SignUpResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}}
	var (
		user   = new(po.User)
		avatar *po.Avatar
		err    error
	)
	copier.Copy(user, req)
	user.AvatarKey = constant.CONST_AVATAR_KEY_SMALL
	user.Password = utils.MD5(user.Password)
	user.ServerId = utils.NewServerId(0, req.ServerId, req.RegPlatform)
	user.LarkId = xsnowflake.DefaultLarkId()
	user.Hash = utils.MD5(user.LarkId)

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		// mobile 重复校验
		err = s.RecheckMobile(tx, -1, req.Mobile, resp)
		if err != nil {
			return
		}
		err = s.authRepo.TxCreate(tx, user)
		if err != nil {
			resp.Set(ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR)
			return
		}
		avatar = &po.Avatar{
			OwnerId:      user.Uid,
			OwnerType:    int(pb_enum.AVATAR_OWNER_USER_AVATAR),
			AvatarSmall:  constant.CONST_AVATAR_KEY_SMALL,
			AvatarMedium: constant.CONST_AVATAR_KEY_MEDIUM,
			AvatarLarge:  constant.CONST_AVATAR_KEY_LARGE,
		}
		err = s.avatarRepo.TxCreate(tx, avatar)
		if err != nil {
			resp.Set(ERROR_CODE_AUTH_INSERT_VALUE_FAILED, ERROR_AUTH_INSERT_VALUE_FAILED)
			return
		}
		return
	})
	if err != nil {
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	var (
		accessToken  *pb_auth.Token
		refreshToken *pb_auth.Token
	)
	accessToken, refreshToken, err = s.CreateToken(user.Uid, req.RegPlatform)
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
		terr := s.userCache.SetUserAndServer(resp.UserInfo, user.ServerId)
		if terr != nil {
			xlog.Warn(terr.Error())
		}
	})
	return
}

func (s *authService) CreateToken(uid int64, platform pb_enum.PLATFORM_TYPE) (aToken *pb_auth.Token, rToken *pb_auth.Token, err error) {
	var (
		accessToken  *xjwt.JwtToken
		refreshToken *xjwt.JwtToken
	)
	accessToken, err = xjwt.CreateToken(uid, int32(platform), true, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	if err != nil {
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, err.Error())
		return
	}
	refreshToken, err = xjwt.CreateToken(uid, int32(platform), false, constant.CONST_DURATION_SHA_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND)
	if err != nil {
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, err.Error())
		return
	}
	err = s.authCache.SetSessionId(uid, int32(platform), accessToken.SessionId, refreshToken.SessionId)
	if err != nil {
		xlog.Warn(ERROR_CODE_AUTH_REDIS_SET_FAILED, ERROR_AUTH_REDIS_SET_FAILED, err.Error())
		return
	}
	aToken = &pb_auth.Token{
		Token:  accessToken.Token,
		Expire: accessToken.Expire,
	}
	rToken = &pb_auth.Token{
		Token:  refreshToken.Token,
		Expire: refreshToken.Expire,
	}
	return
}

func (s *authService) RecheckMobile(tx *gorm.DB, uid int64, mobile string, resp *pb_auth.SignUpResp) (err error) {
	var (
		w      = entity.NewMysqlWhere()
		exists bool
	)
	w.SetFilter("mobile=?", mobile)
	exists, err = s.userRepo.TxExists(tx, w, uid)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED)
		return
	}
	if exists == true {
		err = ERR_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT
		resp.Set(ERROR_CODE_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT, ERROR_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT)
		return
	}
	return
}
