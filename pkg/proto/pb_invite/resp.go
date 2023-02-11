package pb_invite

func (r *InitiateChatInviteResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *ChatInviteListResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *ChatInviteHandleResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
