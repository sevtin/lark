package xoauth2

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"lark/pkg/conf"
)

func NewGoogleOauthConfig(cfg *conf.GoogleOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectUrl,
		Scopes:       cfg.Scopes,
		Endpoint:     google.Endpoint,
	}
}
