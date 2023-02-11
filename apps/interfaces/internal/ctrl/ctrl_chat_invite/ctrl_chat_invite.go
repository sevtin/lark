package ctrl_chat_invite

import (
	"lark/apps/interfaces/internal/service/svc_chat_invite"
)

type ChatInviteCtrl struct {
	chatInviteService svc_chat_invite.ChatInviteService
}

func NewChatInviteCtrl(chatInviteService svc_chat_invite.ChatInviteService) *ChatInviteCtrl {
	return &ChatInviteCtrl{chatInviteService: chatInviteService}
}
