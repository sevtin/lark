package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat"
)

func provideChat() {
	container.Provide(svc_chat.NewChatService)
}
