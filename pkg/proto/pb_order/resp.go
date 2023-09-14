package pb_order

func (r *CreateRedEnvelopeOrderResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
