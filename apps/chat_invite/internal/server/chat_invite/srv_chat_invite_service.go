package chat_invite

import (
	"context"
	"lark/pkg/proto/pb_invite"
)

func (s *chatInviteServer) InitiateChatInvite(ctx context.Context, req *pb_invite.InitiateChatInviteReq) (resp *pb_invite.InitiateChatInviteResp, err error) {
	return s.requestService.InitiateChatInvite(ctx, req)
}

func (s *chatInviteServer) ChatInviteHandle(ctx context.Context, req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp, err error) {
	return s.requestService.ChatInviteHandle(ctx, req)
}

func (s *chatInviteServer) ChatInviteList(ctx context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, err error) {
	return s.requestService.ChatInviteList(ctx, req)
}
