package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat"
)

func provideChat() {
	Provide(svc_chat.NewChatService)
}
