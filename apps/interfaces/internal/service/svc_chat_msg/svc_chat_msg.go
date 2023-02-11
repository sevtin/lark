package svc_chat_msg

import (
	"lark/apps/chat_msg/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	msg_client "lark/apps/message/client"
	"lark/pkg/xhttp"
)

type ChatMessageService interface {
	GetChatMessageList(req *dto_chat_msg.GetChatMessageListReq) (resp *xhttp.Resp)
	// 弃用
	GetChatMessages(req *dto_chat_msg.GetChatMessagesReq) (resp *xhttp.Resp)
	Search(req *dto_chat_msg.SearchMessageReq, uid int64) (resp *xhttp.Resp)
	MessageOperation(req *dto_chat_msg.MessageOperationReq, uid int64, platform int32) (resp *xhttp.Resp)
	SendChatMessage(req *dto_chat_msg.SendChatMessageReq, uid int64, platform int32) (resp *xhttp.Resp)
}

type chatMessageService struct {
	chatMessageClient chat_msg_client.ChatMessageClient
	msgClient         msg_client.MsgClient
}

func NewChatMessageService() ChatMessageService {
	conf := config.GetConfig()
	chatMessageClient := chat_msg_client.NewChatMessageClient(conf.Etcd, conf.ChatMsgServer, conf.Jaeger, conf.Name)
	msgClient := msg_client.NewMsgClient(conf.Etcd, conf.MessageServer, conf.Jaeger, conf.Name)
	return &chatMessageService{chatMessageClient: chatMessageClient, msgClient: msgClient}
}
