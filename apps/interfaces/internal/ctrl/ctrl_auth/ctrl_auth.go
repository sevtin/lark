package ctrl_auth

import (
	"golang.org/x/oauth2"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/service/svc_auth"
	"lark/pkg/common/xoauth2"
)

type AuthCtrl struct {
	authService       svc_auth.AuthService
	googleOauthConfig *oauth2.Config
	githubOauthConfig *oauth2.Config
}

func NewAuthCtrl(authService svc_auth.AuthService) *AuthCtrl {
	srv := &AuthCtrl{authService: authService}
	conf := config.GetConfig()
	srv.googleOauthConfig = xoauth2.NewGoogleOauthConfig(conf.Google)
	srv.githubOauthConfig = xoauth2.NewGithubOauthConfig(conf.Github)
	return srv
}
