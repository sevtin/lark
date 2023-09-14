package xgoogle

import (
	"golang.org/x/oauth2"
	"lark/pkg/conf"
)

func NewOauthConfig(cfg *conf.GoogleOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectUrl,
		Scopes:       cfg.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.AuthURL,
			TokenURL: cfg.TokenURL,
		},
	}
}
