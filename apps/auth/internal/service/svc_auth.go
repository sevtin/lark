package service

import (
	"context"
	"lark/apps/auth/internal/config"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/proto/pb_auth"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error)
	SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
	SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error)
}

type authService struct {
	cfg            *config.Config
	authRepo       repo.AuthRepository
	userRepo       repo.UserRepository
	avatarRepo     repo.AvatarRepository
	chatMemberRepo repo.ChatMemberRepository
	authCache      cache.AuthCache
	userCache      cache.UserCache
}

func NewAuthService(cfg *config.Config,
	authRepo repo.AuthRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	authCache cache.AuthCache,
	userCache cache.UserCache) AuthService {
	return &authService{cfg: cfg,
		authRepo:       authRepo,
		userRepo:       userRepo,
		avatarRepo:     avatarRepo,
		chatMemberRepo: chatMemberRepo,
		authCache:      authCache,
		userCache:      userCache}
}
