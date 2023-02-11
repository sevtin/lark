package dto_chat

type GroupChatDetailsReq struct {
	ChatId int64 `form:"chat_id" json:"chat_id" validate:"required,gt=0"`
}
