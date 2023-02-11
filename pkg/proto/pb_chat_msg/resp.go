package pb_chat_msg

func (r *GetChatMessagesResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *GetChatMessageListResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *SearchMessageResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *MessageOperationResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
