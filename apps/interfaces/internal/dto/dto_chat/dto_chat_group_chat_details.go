package dto_chat

type GroupChatDetailsReq struct {
	ChatId int64 `form:"chat_id" json:"chat_id" binding:"required,gt=0"`
}
