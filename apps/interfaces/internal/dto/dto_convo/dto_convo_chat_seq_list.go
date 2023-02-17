package dto_convo

type ConvoChatSeqListReq struct {
	ChatIds   string `form:"chat_ids" json:"chat_ids" validate:"required,min=19,max=400"`
	Timestamp int64  `form:"timestamp" json:"timestamp" validate:"omitempty,gte=0"`
}
