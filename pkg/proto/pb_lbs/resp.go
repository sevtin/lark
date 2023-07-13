package pb_lbs

func (r *ReportLngLatResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *PeopleNearbyResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
