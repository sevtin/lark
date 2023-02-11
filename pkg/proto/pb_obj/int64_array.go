package pb_obj

import (
	"lark/pkg/utils"
	"strings"
)

func MemberInt64Array(str string) (array *Int64Array, serverId int64) {
	var (
		arr []string
	)
	arr = strings.Split(str, ",")
	if len(arr) != 3 {
		return
	}
	array = &Int64Array{Vals: make([]int64, 4)}

	// 0:ServerId, 1:Platform, 2:Uid, 3:Status
	serverId, _ = utils.ToInt64(arr[0])
	array.Vals[2], _ = utils.ToInt64(arr[1])
	array.Vals[3], _ = utils.ToInt64(arr[2])
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
