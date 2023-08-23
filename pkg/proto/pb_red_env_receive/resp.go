package pb_red_env_receive

func (r *GrabRedEnvelopeResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *OpenRedEnvelopeResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
