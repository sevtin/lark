package dto_chat_member

type ChatMemberListReq struct {
	ChatId  int64 `form:"chat_id" json:"chat_id" validate:"required,gt=0"`
	Limit   int32 `form:"limit" json:"limit" validate:"required,gte=0,lte=200"`
	LastUid int32 `form:"last_uid" json:"last_uid" validate:"omitempty,gte=0"`
	//Page    int32 `form:"page" json:"page" validate:"required,gte=1"`
}
