package service

import (
	"github.com/jinzhu/copier"
	"lark/domain/do"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) oauth2Logic(user *po.OauthUser, platform pb_enum.PLATFORM_TYPE) (aui *pb_auth.AuthUserInfo, err error) {
	aui = &pb_auth.AuthUserInfo{
		UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}},
	}
	var (
		q         = entity.NewMysqlQuery()
		uid       int64
		oauthUser *pdo.OauthUser
		server    *pb_auth.ServerInfo
	)
	q.SetFilter("channel=?", user.Channel)
	q.SetFilter("openid=?", user.Openid)
	oauthUser, err = s.oauthUserRepo.GetOAuthUser(q)
	if err != nil {
		aui.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_QUERY_FAILED, ERROR_AUTH_OAUTH_USER_INFO_QUERY_FAILED)
		return
	}
	server = s.getWsServer()
	if oauthUser.Uid > 0 {
		//已经注册
		err = s.updateGithubUserInfo(user)
		if err != nil {
			aui.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_QUERY_FAILED, ERROR_AUTH_OAUTH_USER_INFO_QUERY_FAILED)
			return
		}
		var (
			signIn *do.SignIn
		)
		q.Normal()
		q.SetFilter("uid=?", oauthUser.Uid)
		signIn = s.signInTransaction(q, platform)
		if signIn.Err != nil || signIn.Code > 0 {
			aui.Set(signIn.Code, signIn.Msg)
			return
		}
		uid = oauthUser.Uid
		_ = copier.Copy(aui.UserInfo, signIn.User)
		_ = copier.Copy(aui.UserInfo.Avatar, signIn.Avatar)
		aui.AccessToken = signIn.AccessToken
		aui.RefreshToken = signIn.RefreshToken

		_, _, err = s.online.UserOnline(uid, int64(server.ServerId), platform)
		if err != nil {
			aui.Set(ERROR_CODE_AUTH_UPDATE_USER_SERVER_ID_FAILED, ERROR_AUTH_UPDATE_USER_SERVER_ID_FAILED)
			xlog.Warn(aui.Code, aui.Msg, err.Error())
			return
		}
	} else {
		//首次注册
		var (
			srvId  int64
			signUp *do.SignUp
		)
		srvId, signUp, err = s.registerUser(user, platform, int64(server.ServerId))
		if signUp.Err != nil || signUp.Code > 0 {
			aui.Set(signUp.Code, signUp.Msg)
			return
		}
		uid = signUp.User.Uid
		_ = copier.Copy(aui.UserInfo, signUp.User)
		_ = copier.Copy(aui.UserInfo.Avatar, signUp.Avatar)
		aui.AccessToken = signUp.AccessToken
		aui.RefreshToken = signUp.RefreshToken

		err = s.userCache.SetServerId(user.Uid, srvId)
		if err != nil {
			aui.Set(ERROR_CODE_AUTH_UPDATE_USER_SERVER_ID_FAILED, ERROR_AUTH_UPDATE_USER_SERVER_ID_FAILED)
			xlog.Warn(aui.Code, aui.Msg, err.Error())
			return
		}
	}
	aui.Server = server
	return
}

func (s *authService) updateGithubUserInfo(user *po.OauthUser) (err error) {
	var (
		u = entity.NewMysqlUpdate()
	)
	u.SetFilter("channel=?", user.Channel)
	u.SetFilter("openid=?", user.Openid)
	u.Set("access_token", user.AccessToken)
	if user.Expire > 0 {
		u.Set("expire", user.Expire)
	}
	if user.RefreshToken != "" {
		u.Set("refresh_token", user.RefreshToken)
	}
	if user.Scope != "" {
		u.Set("scope", user.Scope)
	}
	err = s.oauthUserRepo.UpdateOauthUser(u)
	return
}

func (s *authService) registerUser(oauthUser *po.OauthUser, platform pb_enum.PLATFORM_TYPE, serverId int64) (srvId int64, signUp *do.SignUp, err error) {
	var (
		user   *po.User
		avatar *po.Avatar
	)
	srvId = utils.NewServerId(0, serverId, platform)
	oauthUser.Uid = xsnowflake.NewSnowflakeID()
	oauthUser.OauthId = xsnowflake.NewSnowflakeID()
	user = &po.User{
		Uid:         oauthUser.Uid,
		LarkId:      xsnowflake.DefaultLarkId(),
		Password:    DEFAULT_LOGIN_PASSWORD,
		Udid:        utils.NewUUID(),
		Status:      0,
		Nickname:    oauthUser.Nickname,
		Firstname:   "",
		Lastname:    "",
		Gender:      0,
		BirthTs:     0,
		Email:       oauthUser.Email,
		Mobile:      "",
		RegPlatform: int(platform),
		ServerId:    srvId,
		CityId:      0,
		Avatar:      oauthUser.AvatarUrl,
	}
	avatar = &po.Avatar{
		OwnerId:      user.Uid,
		OwnerType:    int(pb_enum.AVATAR_OWNER_USER_AVATAR),
		AvatarSmall:  oauthUser.AvatarUrl,
		AvatarMedium: oauthUser.AvatarUrl,
		AvatarLarge:  oauthUser.AvatarUrl,
	}
	tx := xmysql.GetTX()
	err = s.oauthUserRepo.TxCreateOauthUser(tx, oauthUser)
	if err != nil {
		tx.Rollback()
	}
	signUp = s.signUpTransaction(tx, user, avatar, platform)
	return
}
