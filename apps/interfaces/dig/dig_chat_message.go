package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat_msg"
)

func provideChatMessage() {
	container.Provide(svc_chat_msg.NewChatMessageService)
}
