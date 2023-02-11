package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat_member"
)

func provideChatMember() {
	container.Provide(svc_chat_member.NewChatMemberService)
}
