package ctrl_chat_member

import (
	"lark/apps/interfaces/internal/service/svc_chat_member"
)

type ChatMemberCtrl struct {
	chatMemberService svc_chat_member.ChatMemberService
}

func NewChatMemberCtrl(chatMemberService svc_chat_member.ChatMemberService) *ChatMemberCtrl {
	return &ChatMemberCtrl{chatMemberService: chatMemberService}
}
