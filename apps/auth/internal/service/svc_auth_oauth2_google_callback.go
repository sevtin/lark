package service

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/xhttp"
)

/*
https://console.cloud.google.com/apis/dashboard
http://localhost:8088/open/auth/google/callback
http://localhost:8088/open/auth/google/auth_code_url
*/
func (s *authService) GoogleOAuth2Callback(ctx context.Context, req *pb_auth.GoogleOAuth2CallbackReq) (resp *pb_auth.GoogleOAuth2CallbackResp, _ error) {
	resp = new(pb_auth.GoogleOAuth2CallbackResp)
	var (
		token *oauth2.Token
		info  *po.OauthUser
		aui   *pb_auth.AuthUserInfo
		err   error
	)
	token, err = s.googleOauthConfig.Exchange(context.Background(), req.Code)
	if err != nil {
		return
	}
	info, err = s.getGoogleUserInfo(token)
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

func (s *authService) getGoogleUserInfo(token *oauth2.Token) (oauthUser *po.OauthUser, err error) {
	var (
		buf  []byte
		info = new(do.GoogleUserInfo)
	)
	buf, err = xhttp.Get(API_GOOGLE_USERINFO+token.AccessToken, nil)
	if err != nil {
		return
	}
	if len(buf) == 0 {
		return
	}
	err = json.Unmarshal(buf, info)
	if err != nil || info == nil {
		return
	}
	oauthUser = &po.OauthUser{
		OauthId:      0,
		Channel:      int(pb_enum.LOGIN_CHANNEL_GOOGLE),
		Openid:       info.ID,
		Uid:          0,
		Username:     info.Email,
		Nickname:     info.Name,
		Email:        info.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expire:       int(token.Expiry.Unix()),
		AvatarUrl:    info.Picture,
		Scope:        "",
	}
	return
}
