package pb_msg

func (r *SendChatMessageResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *MessageOperationResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
