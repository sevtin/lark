package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"strconv"
)

/*
https://github.com/login/oauth/authorize?client_id=0d913b42b82360544ba4&redirect_uri=http://localhost:8088/open/auth/github/callback&login&scope=repo&state&allow_signup=true
*/
func (s *authService) GithubOAuth2Callback(ctx context.Context, req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp, _ error) {
	resp = &pb_auth.GithubOAuth2CallbackResp{AuthUserInfo: &pb_auth.AuthUserInfo{}}
	var (
		token *oauth2.Token
		info  *po.OauthUser
		aui   *pb_auth.AuthUserInfo
		err   error
	)
	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()
	token, err = s.getToken(req.Code)
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

func (s *authService) getToken(code string) (token *oauth2.Token, err error) {
	/*
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
	*/
	token, err = s.githubOauthConfig.Exchange(context.Background(), code)
	return
}

func (s *authService) getGithubUserInfo(token *oauth2.Token) (oauthUser *po.OauthUser, err error) {
	/*
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
	*/
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var (
		info *github.User
		resp *github.Response
	)
	info, resp, err = client.Users.Get(ctx, "")
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	if info == nil {
		return
	}
	oauthUser = &po.OauthUser{
		OauthId:      0,
		Channel:      int(pb_enum.LOGIN_CHANNEL_GITHUB),
		Openid:       strconv.Itoa(int(info.GetID())),
		Uid:          0,
		Username:     info.GetLogin(),
		Nickname:     info.GetLogin(),
		Email:        info.GetEmail(),
		AccessToken:  token.AccessToken,
		RefreshToken: "",
		Expire:       0,
		AvatarUrl:    info.GetAvatarURL(),
		Url:          info.GetHTMLURL(),
		Scope:        "",
	}
	return
}
