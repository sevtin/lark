package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat_member"
)

func init() {
	Provide(svc_chat_member.NewChatMemberService)
}
