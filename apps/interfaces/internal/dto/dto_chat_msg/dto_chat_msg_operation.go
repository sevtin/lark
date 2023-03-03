package dto_chat_msg

type MessageOperationReq struct {
	ChatId int64 `json:"chat_id" binding:"required,gt=0"`
	SeqId  int64 `json:"seq_id" binding:"required,gt=0"`
	Opn    int32 `json:"opn" binding:"required,gte=1,lte=127"`
}
