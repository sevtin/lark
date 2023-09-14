package service

import (
	"context"
	"golang.org/x/oauth2"
	"lark/apps/auth/internal/config"
	chat_member_client "lark/apps/chat_member/client"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xgoogle"
	"lark/pkg/proto/pb_auth"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error)
	SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
	SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error)
	GithubOAuth2Callback(ctx context.Context, req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp, err error)
	GoogleOAuth2Callback(ctx context.Context, req *pb_auth.GoogleOAuth2CallbackReq) (resp *pb_auth.GoogleOAuth2CallbackResp, err error)
}

type authService struct {
	cfg               *config.Config
	authRepo          repo.AuthRepository
	userRepo          repo.UserRepository
	avatarRepo        repo.AvatarRepository
	chatMemberRepo    repo.ChatMemberRepository
	authCache         cache.AuthCache
	userCache         cache.UserCache
	svrMgrCache       cache.ServerMgrCache
	chatMemberClient  chat_member_client.ChatMemberClient
	googleOauthConfig *oauth2.Config
}

func NewAuthService(cfg *config.Config,
	authRepo repo.AuthRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	authCache cache.AuthCache,
	userCache cache.UserCache,
	svrMgrCache cache.ServerMgrCache) AuthService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.Jaeger, cfg.Name)
	svc := &authService{cfg: cfg,
		authRepo:         authRepo,
		userRepo:         userRepo,
		avatarRepo:       avatarRepo,
		chatMemberRepo:   chatMemberRepo,
		authCache:        authCache,
		userCache:        userCache,
		svrMgrCache:      svrMgrCache,
		chatMemberClient: chatMemberClient,
	}
	svc.googleOauthConfig = xgoogle.NewOauthConfig(cfg.Google)
	return svc
}
