package chat_msg

import (
	"context"
	"lark/pkg/proto/pb_chat_msg"
)

func (s *chatMessageServer) GetChatMessageList(ctx context.Context, req *pb_chat_msg.GetChatMessageListReq) (resp *pb_chat_msg.GetChatMessageListResp, err error) {
	return s.chatMessageService.GetChatMessageList(ctx, req)
}

// 弃用
//func (s *chatMessageServer) GetChatMessages(ctx context.Context, req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp, err error) {
//	return s.chatMessageService.GetChatMessages(ctx, req)
//}

func (s *chatMessageServer) SearchMessage(ctx context.Context, req *pb_chat_msg.SearchMessageReq) (resp *pb_chat_msg.SearchMessageResp, err error) {
	return s.chatMessageService.SearchMessage(ctx, req)
}

func (s *chatMessageServer) MessageOperation(ctx context.Context, req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp, err error) {
	return s.chatMessageService.MessageOperation(ctx, req)
}
