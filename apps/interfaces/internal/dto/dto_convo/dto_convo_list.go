package dto_convo

type ConvoListReq struct {
	ChatIds string `form:"chat_ids" json:"chat_ids" binding:"required,min=19,max=400"`
}
