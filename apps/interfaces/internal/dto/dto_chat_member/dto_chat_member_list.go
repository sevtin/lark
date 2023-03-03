package dto_chat_member

type ChatMemberListReq struct {
	ChatId  int64 `form:"chat_id" json:"chat_id" binding:"required,gt=0"`
	Limit   int32 `form:"limit" json:"limit" binding:"required,gte=0,lte=200"`
	LastUid int32 `form:"last_uid" json:"last_uid" binding:"omitempty,gte=0"`
	//Page    int32 `form:"page" json:"page" binding:"required,gte=1"`
}
