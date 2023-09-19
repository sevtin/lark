package service

import (
	"context"
	"fmt"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"strconv"
)

/*
https://github.com/login/oauth/authorize?client_id=Iv1.4a3ecfaf2c5e4f1a&redirect_uri=http://localhost:8088/open/auth/github/callback&login&scope=repo&state&allow_signup=true
*/
func (s *authService) GithubOAuth2Callback(ctx context.Context, req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp, _ error) {
	resp = &pb_auth.GithubOAuth2CallbackResp{AuthUserInfo: &pb_auth.AuthUserInfo{}}
	var (
		token *do.GithubToken
		info  *po.OauthUser
		aui   *pb_auth.AuthUserInfo
		err   error
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
	info, err = s.getGithubUserInfo(token)
	if err != nil || info == nil {
		resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED)
		return
	}
	if info.Openid == "" {
		resp.Set(ERROR_CODE_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED, ERROR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED)
		err = ERR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED
		return
	}
	aui, err = s.oauth2Logic(info, req.Platform)
	if err != nil || aui == nil {
		return
	}
	resp.Set(aui.Code, aui.Msg)
	resp.AuthUserInfo = aui
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

func (s *authService) getGithubUserInfo(token *do.GithubToken) (oauthUser *po.OauthUser, err error) {
	var (
		buf  []byte
		info *do.GithubUserInfo
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
	info = new(do.GithubUserInfo)
	err = utils.Unmarshal(string(buf), info)
	if err != nil {
		return
	}
	oauthUser = &po.OauthUser{
		OauthId:      0,
		Channel:      int(pb_enum.LOGIN_CHANNEL_GITHUB),
		Openid:       strconv.Itoa(info.ID),
		Uid:          0,
		Username:     info.Login,
		Nickname:     info.Login,
		Email:        info.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: "",
		Expire:       0,
		AvatarUrl:    info.AvatarURL,
		Scope:        token.Scope,
	}
	return
}
