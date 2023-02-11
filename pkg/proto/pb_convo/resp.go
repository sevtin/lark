package pb_convo

func (r *ConvoListResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *ConvoChatSeqListResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
