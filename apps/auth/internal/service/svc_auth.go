package service

import (
	"context"
	"golang.org/x/oauth2"
	"lark/apps/auth/internal/config"
	chat_member_client "lark/apps/chat_member/client"
	"lark/business/biz_online"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xoauth2"
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
	oauthUserRepo     repo.OauthUserRepository
	userRepo          repo.UserRepository
	avatarRepo        repo.AvatarRepository
	chatMemberRepo    repo.ChatMemberRepository
	authCache         cache.AuthCache
	userCache         cache.UserCache
	svrMgrCache       cache.ServerMgrCache
	chatMemberClient  chat_member_client.ChatMemberClient
	googleOauthConfig *oauth2.Config
	githubOauthConfig *oauth2.Config
	online            biz_online.Online
}

func NewAuthService(cfg *config.Config,
	oauthUserRepo repo.OauthUserRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	authCache cache.AuthCache,
	userCache cache.UserCache,
	svrMgrCache cache.ServerMgrCache,
	online biz_online.Online) AuthService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.Jaeger, cfg.Name)
	svc := &authService{cfg: cfg,
		oauthUserRepo:    oauthUserRepo,
		userRepo:         userRepo,
		avatarRepo:       avatarRepo,
		chatMemberRepo:   chatMemberRepo,
		authCache:        authCache,
		userCache:        userCache,
		svrMgrCache:      svrMgrCache,
		chatMemberClient: chatMemberClient,
		online:           online,
	}
	svc.googleOauthConfig = xoauth2.NewGoogleOauthConfig(cfg.Google)
	svc.githubOauthConfig = xoauth2.NewGithubOauthConfig(cfg.Github)
	return svc
}
