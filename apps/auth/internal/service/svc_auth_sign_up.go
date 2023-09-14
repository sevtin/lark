package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, _ error) {
	resp = &pb_auth.SignUpResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}}
	var (
		server *pb_auth.ServerInfo
		user   = new(po.User)
		avatar *po.Avatar
		err    error
		signUp *do.SignUp
	)
	server = s.getWsServer()
	_ = copier.Copy(user, req)
	user.Avatar = constant.CONST_AVATAR_SMALL
	user.Password = utils.MD5(user.Password)
	user.ServerId = utils.NewServerId(0, int64(server.ServerId), req.RegPlatform)
	user.Uid = xsnowflake.NewSnowflakeID()
	user.LarkId = xsnowflake.DefaultLarkId()

	avatar = &po.Avatar{
		OwnerId:      user.Uid,
		OwnerType:    int(pb_enum.AVATAR_OWNER_USER_AVATAR),
		AvatarSmall:  constant.CONST_AVATAR_SMALL,
		AvatarMedium: constant.CONST_AVATAR_MEDIUM,
		AvatarLarge:  constant.CONST_AVATAR_LARGE,
	}

	err = s.RecheckMobile(-1, req.Mobile, resp)
	if err != nil {
		return
	}
	signUp = s.signUpTransaction(xmysql.GetTX(), user, avatar, req.RegPlatform)
	if signUp.Err != nil || signUp.Code > 0 {
		resp.Set(signUp.Code, signUp.Msg)
		return
	}

	_ = copier.Copy(resp.UserInfo, user)
	_ = copier.Copy(resp.UserInfo.Avatar, avatar)
	resp.AccessToken = signUp.AccessToken
	resp.RefreshToken = signUp.RefreshToken
	resp.Server = server

	_ = xants.Submit(func() {
		terr := s.userCache.SetUserServer(user.Uid, user.ServerId)
		if terr != nil {
			xlog.Warn(terr.Error())
		}
	})
	return
}

func (s *authService) signUpTransaction(tx *gorm.DB, user *po.User, avatar *po.Avatar, regPlatform pb_enum.PLATFORM_TYPE) (signUp *do.SignUp) {
	signUp = new(do.SignUp)
	signUp.User = user
	signUp.Avatar = avatar
	signUp.Err = s.authRepo.TxCreate(tx, user)
	if signUp.Err != nil {
		signUp.Code = ERROR_CODE_AUTH_REGISTER_ERR
		signUp.Msg = ERROR_AUTH_REGISTER_TYPE_ERR
		tx.Rollback()
		return
	}
	signUp.Err = s.avatarRepo.TxCreate(tx, avatar)
	if signUp.Err != nil {
		signUp.Code = ERROR_CODE_AUTH_INSERT_VALUE_FAILED
		signUp.Msg = ERROR_AUTH_INSERT_VALUE_FAILED
		tx.Rollback()
		return
	}
	tx.Commit()
	signUp.AccessToken, signUp.RefreshToken, signUp.Err = s.createToken(user.Uid, regPlatform)
	if signUp.Err != nil {
		signUp.Code = ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED
		signUp.Msg = ERROR_AUTH_GENERATE_TOKEN_FAILED
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, signUp.Err.Error())
		return
	}
	return
}

func (s *authService) createToken(uid int64, platform pb_enum.PLATFORM_TYPE) (aToken *pb_auth.Token, rToken *pb_auth.Token, err error) {
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

func (s *authService) RecheckMobile(uid int64, mobile string, resp *pb_auth.SignUpResp) (err error) {
	var (
		w      = entity.NewMysqlQuery()
		exists bool
	)
	w.SetFilter("mobile=?", mobile)
	exists, err = s.userRepo.Exists(w, uid)
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

func (s *authService) chatMemberOnOffLine(uid int64, serverId int64, platform pb_enum.PLATFORM_TYPE) (resp *pb_chat_member.ChatMemberOnOffLineResp) {
	req := &pb_chat_member.ChatMemberOnOffLineReq{
		Uid:      uid,
		ServerId: serverId,
		Platform: platform,
	}
	return s.chatMemberClient.ChatMemberOnOffLine(req)
}

func (s *authService) getWsServer() (wsServer *pb_auth.ServerInfo) {
	var (
		member   string
		server   string
		serverId int64
		port     int
	)
	list := s.svrMgrCache.ZRevRangeMsgGateway(0, 0)
	if len(list) > 0 {
		member = list[0]
	}
	server, serverId, port = utils.GetMsgGatewayServer(member)
	wsServer = &pb_auth.ServerInfo{
		ServerId: int32(serverId),
		Name:     server,
		Port:     int32(port),
	}
	return
}
