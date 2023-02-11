package service

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
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
		tx     *gorm.DB
		err    error
	)
	copier.Copy(user, req)
	user.AvatarKey = constant.CONST_AVATAR_KEY_SMALL
	user.Password = utils.MD5(user.Password)
	user.ServerId = utils.NewServerId(0, req.ServerId, req.RegPlatform)

	tx = xmysql.GetTX()
	err = s.authRepo.TxCreate(tx, user)
	if err != nil {
		// 回滚
		tx.Rollback()
		switch err.(type) {
		case *mysql.MySQLError:
			if err.(*mysql.MySQLError).Number == constant.ERROR_CODE_MYSQL_DUPLICATE_ENTRY {
				resp.Set(ERROR_CODE_AUTH_MOBILE_HAS_BEEN_REGISTERED, ERROR_AUTH_MOBILE_HAS_BEEN_REGISTERED)
				xlog.Warn(ERROR_CODE_AUTH_MOBILE_HAS_BEEN_REGISTERED, ERROR_AUTH_MOBILE_HAS_BEEN_REGISTERED, err.Error())
				return
			}
		}
		resp.Set(ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR)
		xlog.Warn(ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR, err.Error())
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
		// 回滚
		tx.Rollback()
		resp.Set(ERROR_CODE_AUTH_INSERT_VALUE_FAILED, ERROR_AUTH_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_INSERT_VALUE_FAILED, ERROR_AUTH_INSERT_VALUE_FAILED, err.Error())
		return
	}
	// 提交
	tx.Commit()
	copier.Copy(resp.UserInfo, user)
	copier.Copy(resp.UserInfo.Avatar, avatar)

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
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	xants.Submit(func() {
		s.userCache.SetUserAndServer(s.cfg.Redis.Prefix, resp.UserInfo, user.ServerId)
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
	err = s.authCache.SetSessionId(s.cfg.Redis.Prefix, uid, int32(platform), accessToken.SessionId, refreshToken.SessionId)
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
