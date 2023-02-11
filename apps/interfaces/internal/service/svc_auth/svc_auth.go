package svc_auth

import (
	auth_client "lark/apps/auth/client"
	chat_member_client "lark/apps/chat_member/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/domain/cache"
	"lark/pkg/xhttp"
)

type AuthService interface {
	SignUp(params *dto_auth.SignUpReq) (resp *xhttp.Resp)
	SignIn(params *dto_auth.SignInReq) (resp *xhttp.Resp)
	RefreshToken(params *dto_auth.RefreshTokenReq) (resp *xhttp.Resp)
	SignOut(params *dto_auth.SignOutReq) (resp *xhttp.Resp)
}

type authService struct {
	authClient       auth_client.AuthClient
	chatMemberClient chat_member_client.ChatMemberClient
	serverMgrCache   cache.ServerMgrCache
}

func NewAuthService(serverMgrCache cache.ServerMgrCache) AuthService {
	conf := config.GetConfig()
	authClient := auth_client.NewAuthClient(conf.Etcd, conf.AuthServer, conf.Jaeger, conf.Name)
	chatMemberClient := chat_member_client.NewChatMemberClient(conf.Etcd, conf.ChatMemberServer, conf.Jaeger, conf.Name)
	return &authService{
		authClient:       authClient,
		chatMemberClient: chatMemberClient,
		serverMgrCache:   serverMgrCache,
	}
}
