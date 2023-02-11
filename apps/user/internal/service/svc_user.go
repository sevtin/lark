package service

import (
	"context"
	"lark/apps/user/internal/config"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/proto/pb_user"
)

type UserService interface {
	GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error)
	GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, err error)
	GetBasicUserInfo(ctx context.Context, req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp, err error)
	EditUserInfo(ctx context.Context, req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp, err error)
	SearchUser(ctx context.Context, req *pb_user.SearchUserReq) (resp *pb_user.SearchUserResp, err error)
	UploadAvatar(ctx context.Context, req *pb_user.UploadAvatarReq) (resp *pb_user.UploadAvatarResp, err error)
	GetBasicUserInfoList(ctx context.Context, req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp, err error)
	GetServerIdList(ctx context.Context, req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp, err error)
}

type userService struct {
	cfg             *config.Config
	userRepo        repo.UserRepository
	avatarRepo      repo.AvatarRepository
	chatMemberRepo  repo.ChatMemberRepository
	userCache       cache.UserCache
	chatMemberCache cache.ChatMemberCache
}

func NewUserService(
	cfg *config.Config,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	userCache cache.UserCache,
	chatMemberCache cache.ChatMemberCache) UserService {
	svc := &userService{
		cfg:             cfg,
		userRepo:        userRepo,
		avatarRepo:      avatarRepo,
		chatMemberRepo:  chatMemberRepo,
		userCache:       userCache,
		chatMemberCache: chatMemberCache,
	}
	return svc
}
