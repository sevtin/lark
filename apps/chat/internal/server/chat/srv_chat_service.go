package chat

import (
	"context"
	"lark/pkg/proto/pb_chat"
)

func (s *chatServer) CreateGroupChat(ctx context.Context, req *pb_chat.CreateGroupChatReq) (resp *pb_chat.CreateGroupChatResp, err error) {
	return s.ChatService.CreateGroupChat(ctx, req)
}

func (s *chatServer) EditGroupChat(ctx context.Context, req *pb_chat.EditGroupChatReq) (resp *pb_chat.EditGroupChatResp, err error) {
	return s.ChatService.EditGroupChat(ctx, req)
}

func (s *chatServer) GroupChatDetails(ctx context.Context, req *pb_chat.GroupChatDetailsReq) (resp *pb_chat.GroupChatDetailsResp, err error) {
	return s.ChatService.GroupChatDetails(ctx, req)
}

func (s *chatServer) RemoveGroupChatMember(ctx context.Context, req *pb_chat.RemoveGroupChatMemberReq) (resp *pb_chat.RemoveGroupChatMemberResp, err error) {
	return s.ChatService.RemoveGroupChatMember(ctx, req)
}

func (s *chatServer) QuitGroupChat(ctx context.Context, req *pb_chat.QuitGroupChatReq) (resp *pb_chat.QuitGroupChatResp, err error) {
	return s.ChatService.QuitGroupChat(ctx, req)
}

func (s *chatServer) DeleteContact(ctx context.Context, req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp, err error) {
	return s.ChatService.DeleteContact(ctx, req)
}

func (s *chatServer) UploadAvatar(ctx context.Context, req *pb_chat.UploadAvatarReq) (resp *pb_chat.UploadAvatarResp, err error) {
	return s.ChatService.UploadAvatar(ctx, req)
}

func (s *chatServer) GetChatInfo(ctx context.Context, req *pb_chat.GetChatInfoReq) (resp *pb_chat.GetChatInfoResp, err error) {
	return s.ChatService.GetChatInfo(ctx, req)
}
