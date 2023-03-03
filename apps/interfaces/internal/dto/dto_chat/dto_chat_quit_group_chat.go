package dto_chat

type QuitGroupChatReq struct {
	ChatId int64 `json:"chat_id" binding:"required,gt=0"`
}
