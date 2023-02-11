package pb_dist

func (r *DistMessageResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
