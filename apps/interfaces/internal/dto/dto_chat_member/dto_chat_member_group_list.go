package dto_chat_member

type GroupChatListReq struct {
	Limit      int32 `form:"limit" json:"limit" validate:"required,gte=10,lte=100"`
	LastChatId int64 `form:"last_chat_id" json:"last_chat_id" validate:"omitempty,gte=0"`
}
