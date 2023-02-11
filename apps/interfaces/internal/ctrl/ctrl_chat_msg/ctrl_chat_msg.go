package ctrl_chat_msg

import (
	"lark/apps/interfaces/internal/service/svc_chat_msg"
)

type ChatMessageCtrl struct {
	chatMessageService svc_chat_msg.ChatMessageService
}

func NewChatMessageCtrl(chatMessageService svc_chat_msg.ChatMessageService) *ChatMessageCtrl {
	return &ChatMessageCtrl{chatMessageService: chatMessageService}
}
