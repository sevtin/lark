package pb_pay

func (r *AlipayReturnResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *AlipayNotifyResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
