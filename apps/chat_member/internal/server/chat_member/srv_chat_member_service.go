package chat_member

import (
	"context"
	"lark/pkg/proto/pb_chat_member"
)

func (s *chatMemberServer) GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, err error) {
	return s.chatMemberService.GetChatMemberInfo(ctx, req)
}

func (s *chatMemberServer) ChatMemberOnOffLine(ctx context.Context, req *pb_chat_member.ChatMemberOnOffLineReq) (resp *pb_chat_member.ChatMemberOnOffLineResp, err error) {
	return s.chatMemberService.ChatMemberOnOffLine(ctx, req)
}

func (s *chatMemberServer) GetDistMemberList(ctx context.Context, req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp, err error) {
	return s.chatMemberService.GetDistMemberList(ctx, req)
}

func (s *chatMemberServer) GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, err error) {
	return s.chatMemberService.GetChatMemberList(ctx, req)
}

func (s *chatMemberServer) GetContactList(ctx context.Context, req *pb_chat_member.GetContactListReq) (resp *pb_chat_member.GetContactListResp, err error) {
	return s.chatMemberService.GetContactList(ctx, req)
}

func (s *chatMemberServer) GetGroupChatList(ctx context.Context, req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp, err error) {
	return s.chatMemberService.GetGroupChatList(ctx, req)
}
