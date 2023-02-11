package dto_chat

type QuitGroupChatReq struct {
	ChatId int64 `json:"chat_id" validate:"required,gt=0"`
}
