package dto_chat_member

type ContactListReq struct {
	Limit      int32 `form:"limit" json:"limit" binding:"required,gte=10,lte=100"`
	LastChatId int64 `form:"last_chat_id" json:"last_chat_id" binding:"omitempty,gte=0"`
}
