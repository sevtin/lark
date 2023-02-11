package service

import (
	"lark/apps/oauth/internal/config"
)

type OAuthService interface {
}

type oauthService struct {
	cfg *config.Config
}

func NewOAuthService(cfg *config.Config) OAuthService {
	return &oauthService{cfg: cfg}
}
