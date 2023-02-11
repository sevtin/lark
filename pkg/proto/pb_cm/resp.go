package pb_cm

func (r *CloudMessageResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
