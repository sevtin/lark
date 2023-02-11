package user

import (
	"context"
	"lark/pkg/proto/pb_user"
)

func (s *userServer) EditUserInfo(ctx context.Context, req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp, err error) {
	return s.userService.EditUserInfo(ctx, req)
}

func (s *userServer) GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	return s.userService.GetUserInfo(ctx, req)
}

func (s *userServer) GetBasicUserInfo(ctx context.Context, req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp, err error) {
	return s.userService.GetBasicUserInfo(ctx, req)
}

func (s *userServer) GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, err error) {
	return s.userService.GetUserList(ctx, req)
}

func (s *userServer) SearchUser(ctx context.Context, req *pb_user.SearchUserReq) (resp *pb_user.SearchUserResp, err error) {
	return s.userService.SearchUser(ctx, req)
}

func (s *userServer) UploadAvatar(ctx context.Context, req *pb_user.UploadAvatarReq) (resp *pb_user.UploadAvatarResp, err error) {
	return s.userService.UploadAvatar(ctx, req)
}

func (s *userServer) GetBasicUserInfoList(ctx context.Context, req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp, err error) {
	return s.userService.GetBasicUserInfoList(ctx, req)
}

func (s *userServer) GetServerIdList(ctx context.Context, req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp, err error) {
	return s.userService.GetServerIdList(ctx, req)
}
