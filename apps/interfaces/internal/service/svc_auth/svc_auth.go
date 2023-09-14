package svc_auth

import (
	auth_client "lark/apps/auth/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/xhttp"
)

type AuthService interface {
	SignUp(params *dto_auth.SignUpReq) (resp *xhttp.Resp)
	SignIn(params *dto_auth.SignInReq) (resp *xhttp.Resp)
	RefreshToken(params *dto_auth.RefreshTokenReq) (resp *xhttp.Resp)
	SignOut(params *dto_auth.SignOutReq) (resp *xhttp.Resp)
	GithubOAuth2Callback(params *dto_auth.GithubOauthCallbackReq) (resp *xhttp.Resp)
	GoogleOAuth2Callback(params *dto_auth.GoogleOauthCallbackReq) (resp *xhttp.Resp)
}

type authService struct {
	authClient auth_client.AuthClient
}

func NewAuthService() AuthService {
	conf := config.GetConfig()
	authClient := auth_client.NewAuthClient(conf.Etcd, conf.AuthServer, conf.Jaeger, conf.Name)
	return &authService{authClient: authClient}
}
