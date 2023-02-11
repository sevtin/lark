package pb_avatar

func (r *SetAvatarResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
