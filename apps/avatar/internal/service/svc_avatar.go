package service

import (
	"context"
	"lark/apps/avatar/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_avatar"
)

// 弃用
type AvatarService interface {
	SetAvatar(ctx context.Context, req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp, err error)
}

type avatarService struct {
	cfg            *config.Config
	avatarRepo     repo.AvatarRepository
	chatMemberRepo repo.ChatMemberRepository
	chatRepo       repo.ChatRepository
	userRepo       repo.UserRepository
}

func NewAvatarService(cfg *config.Config, avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	chatRepo repo.ChatRepository,
	userRepo repo.UserRepository) AvatarService {
	svc := &avatarService{cfg: cfg,
		avatarRepo:     avatarRepo,
		chatMemberRepo: chatMemberRepo,
		chatRepo:       chatRepo,
		userRepo:       userRepo,
	}
	return svc
}
