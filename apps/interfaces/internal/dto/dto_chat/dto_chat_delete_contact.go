package dto_chat

type DeleteContactReq struct {
	ChatId    int64 `json:"chat_id" validate:"required,gt=0"`
	ContactId int64 `json:"contact_id" validate:"required,gt=0"`
}
