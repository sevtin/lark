package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat_invite"
)

func provideChatInvite() {
	container.Provide(svc_chat_invite.NewChatInviteService)
}
