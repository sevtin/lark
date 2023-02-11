package pb_gw

func (r *HealthCheckResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *SendTopicMessageResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
