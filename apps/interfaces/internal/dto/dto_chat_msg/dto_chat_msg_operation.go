package dto_chat_msg

type MessageOperationReq struct {
	ChatId int64 `json:"chat_id" validate:"required,gt=0"`
	SeqId  int64 `json:"seq_id" validate:"required,gt=0"`
	Opn    int32 `json:"opn" validate:"required,gte=1,lte=127"`
}
