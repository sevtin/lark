package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat"
)

func init() {
	Provide(svc_chat.NewChatService)
}
