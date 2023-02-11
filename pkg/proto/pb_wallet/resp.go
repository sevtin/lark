package pb_wallet

func (r *GetBalancesResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *ExchangeResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *TransferResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *RechargeResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *WithdrawResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
