package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"lark/domain/do"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (s *authService) GithubOAuth2Callback(ctx context.Context, req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp, _ error) {
	resp = &pb_auth.GithubOAuth2CallbackResp{UserInfo: &pb_user.UserInfo{}}
	var (
		token     *do.GithubToken
		info      *do.GithubUser
		q         = entity.NewMysqlQuery()
		oauthUser *pdo.OauthUser
		server    *pb_auth.ServerInfo
		err       error
	)
	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()
	token, err = s.getToken(s.getTokenAuthUrl(req.Code))
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED)
		return
	}
	if token.AccessToken == "" {
		resp.Set(ERROR_CODE_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED)
		err = ERR_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED
		return
	}
	info, err = s.getUserInfo(token)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED)
		return
	}
	if info.ID == 0 {
		resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED)
		err = ERR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED
		return
	}
	q.SetFilter("channel=?", pb_enum.LOGIN_CHANNEL_GITHUB)
	q.SetFilter("openid=?", info.ID)
	oauthUser, err = s.authRepo.GetOAuthUser(q)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_QUERY_FAILED, ERROR_AUTH_OAUTH_USER_INFO_QUERY_FAILED)
		return
	}
	server = s.getWsServer()
	if oauthUser.Uid > 0 {
		//已经注册
		err = s.updateGithubUserInfo(info, token)
		if err != nil {
			resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_QUERY_FAILED, ERROR_AUTH_OAUTH_USER_INFO_QUERY_FAILED)
			return
		}
		var (
			signIn    *do.SignIn
			onOffResp *pb_chat_member.ChatMemberOnOffLineResp
		)
		q.Reset()
		q.SetFilter("uid=?", oauthUser.Uid)
		signIn = s.signInTransaction(q, req.Platform)
		if signIn.Err != nil || signIn.Code > 0 {
			resp.Set(signIn.Code, signIn.Msg)
			return
		}
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
	} else {
		//首次注册
		var (
			signUp *do.SignUp
		)
		signUp, err = s.registerGithubUser(info, token, req.Platform, int64(server.ServerId))
		if signUp.Err != nil || signUp.Code > 0 {
			resp.Set(signUp.Code, signUp.Msg)
			return
		}
		copier.Copy(resp.UserInfo, signUp.User)
		copier.Copy(resp.UserInfo.Avatar, signUp.Avatar)
		resp.AccessToken = signUp.AccessToken
		resp.RefreshToken = signUp.RefreshToken
	}
	resp.Server = server
	xants.Submit(func() {
		s.userCache.SetUserInfo(resp.UserInfo)
	})
	return
}

func (s *authService) updateGithubUserInfo(info *do.GithubUser, token *do.GithubToken) (err error) {
	var (
		u = entity.NewMysqlUpdate()
	)
	u.SetFilter("channel=?", pb_enum.LOGIN_CHANNEL_GITHUB)
	u.SetFilter("openid=?", info.ID)
	u.Set("access_token", token.AccessToken)
	if token.Scope != "" {
		u.Set("scope", token.Scope)
	}
	err = s.authRepo.UpdateOauthUser(u)
	return
}

// TODO:头像字段key/url待统一
func (s *authService) registerGithubUser(info *do.GithubUser, token *do.GithubToken, platform pb_enum.PLATFORM_TYPE, serverId int64) (signUp *do.SignUp, err error) {
	var (
		oauthUser *po.OauthUser
		user      *po.User
		avatar    *po.Avatar
		uid       = xsnowflake.NewSnowflakeID()
	)
	oauthUser = &po.OauthUser{
		OauthId:      xsnowflake.NewSnowflakeID(),
		Channel:      int(pb_enum.LOGIN_CHANNEL_GITHUB),
		Openid:       utils.ToString(info.ID),
		Uid:          uid,
		Username:     info.Login,
		Nickname:     info.Name,
		Email:        info.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: "",
		Expire:       0,
		AvatarUrl:    info.AvatarURL,
		Scope:        token.Scope,
	}
	user = &po.User{
		Uid:         uid,
		LarkId:      xsnowflake.DefaultLarkId(),
		Password:    DEFAULT_LOGIN_PASSWORD,
		Udid:        utils.NewUUID(),
		Status:      0,
		Nickname:    info.Login,
		Firstname:   "",
		Lastname:    "",
		Gender:      0,
		BirthTs:     0,
		Email:       info.Email,
		Mobile:      "",
		RegPlatform: int(platform),
		ServerId:    utils.NewServerId(0, serverId, platform),
		CityId:      0,
		Avatar:      info.AvatarURL,
	}
	avatar = &po.Avatar{
		OwnerId:      user.Uid,
		OwnerType:    int(pb_enum.AVATAR_OWNER_USER_AVATAR),
		AvatarSmall:  info.AvatarURL,
		AvatarMedium: info.AvatarURL,
		AvatarLarge:  info.AvatarURL,
	}
	tx := xmysql.GetTX()
	err = s.authRepo.TxCreateOauthUser(tx, oauthUser)
	if err != nil {
		tx.Rollback()
	}
	signUp = s.signUpTransaction(tx, user, avatar, platform)
	return
}

func (s *authService) getTokenAuthUrl(code string) string {
	return fmt.Sprintf(API_GITHUB_OAUTH_ACCESS_TOKEN, s.cfg.Github.ClientId, s.cfg.Github.ClientSecret, code)
}

func (s *authService) getToken(url string) (token *do.GithubToken, err error) {
	var (
		buf []byte
	)
	token = new(do.GithubToken)
	buf, err = xhttp.Get(url, nil, &xhttp.HeaderOption{"accept", "application/json"})
	if err != nil {
		return
	}
	err = utils.Unmarshal(string(buf), token)
	if err != nil {
		return
	}
	return
}

func (s *authService) getUserInfo(token *do.GithubToken) (info *do.GithubUser, err error) {
	var (
		buf []byte
	)
	buf, err = xhttp.Get(API_GITHUB_USER, nil,
		&xhttp.HeaderOption{
			Key:   "accept",
			Value: "application/json"},
		&xhttp.HeaderOption{
			Key:   "Authorization",
			Value: fmt.Sprintf("token %s", token.AccessToken)})
	if err != nil {
		return
	}
	if len(buf) == 0 {
		return
	}
	info = new(do.GithubUser)
	utils.Unmarshal(string(buf), info)
	return
}
