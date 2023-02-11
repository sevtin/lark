package service

import (
	"fmt"
	"lark/apps/oauth/internal/bo"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (s *oauthService) getTokenAuthUrl(code string) string {
	return fmt.Sprintf(API_GITHUB_OAUTH, s.cfg.OAuth.ClientId, s.cfg.OAuth.ClientSecret, code)
}

func (s *oauthService) getToken(url string) (token *bo.Token, err error) {
	var (
		buf []byte
	)
	buf, err = xhttp.Get(url, nil)
	if err != nil {
		return
	}
	token = new(bo.Token)
	utils.Unmarshal(string(buf), token)
	return
}

func (s *oauthService) getUserInfo(token *bo.Token) (info *bo.UserInfo, err error) {
	var (
		buf []byte
	)
	buf, err = xhttp.Post(API_GITHUB_USER, nil,
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
	info = new(bo.UserInfo)
	utils.Unmarshal(string(buf), info)
	return
}
