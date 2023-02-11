package dto_chat

type RemoveGroupChatMemberReq struct {
	ChatId     int64   `json:"chat_id" validate:"required,gt=0"`
	MemberList []int64 `json:"member_list" validate:"required"`
}
