package dto_convo

type ConvoChatSeqListReq struct {
	ChatIds   string `form:"chat_ids" json:"chat_ids" binding:"required,min=19,max=400"`
	Timestamp int64  `form:"timestamp" json:"timestamp" binding:"omitempty,gte=0"`
}
