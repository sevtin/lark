package pb_obj

func MemberInt64Array(uid int64, status int64, srvId int64) (array *Int64Array) {
	array = &Int64Array{Vals: make([]int64, 4)}
	array.Vals[0] = srvId
	array.Vals[2] = uid
	array.Vals[3] = status
	return
}

func (p *Int64Array) GetServerId() int64 {
	return p.Vals[0]
}

func (p *Int64Array) GetPlatform() int64 {
	return p.Vals[1]
}

func (p *Int64Array) GetUid() int64 {
	return p.Vals[2]
}

func (p *Int64Array) GetStatus() int64 {
	return p.Vals[3]
}
