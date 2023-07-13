package logic

import (
	"fmt"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
)

func GetMembersFromHash(hashmap map[string]string) (distMembers map[int64][]*pb_obj.Int64Array) {
	if len(hashmap) == 0 {
		return
	}
	return groupFromHashmap(hashmap)
}

func groupFromHashmap(hashmap map[string]string) (distMembers map[int64][]*pb_obj.Int64Array) {
	var (
		uid      string
		str      string
		array    *pb_obj.Int64Array
		serverId int64
	)
	distMembers = make(map[int64][]*pb_obj.Int64Array)
	for uid, str = range hashmap {
		array, serverId = pb_obj.MemberInt64Array(str, uid)
		setDistMembers(distMembers, serverId, array)
	}
	return
}

func setDistMembers(distMembers map[int64][]*pb_obj.Int64Array, serverId int64, array *pb_obj.Int64Array) {
	if array == nil {
		return
	}
	var (
		iosSid, androidSid, macSid, windowsSid, webSid int64
	)
	iosSid, androidSid, macSid, windowsSid, webSid = utils.GetServerId(serverId)
	if iosSid > 0 {
		putInt64Array(distMembers, array, iosSid, pb_enum.PLATFORM_TYPE_IOS)
	}
	if androidSid > 0 {
		putInt64Array(distMembers, array, androidSid, pb_enum.PLATFORM_TYPE_ANDROID)
	}
	if macSid > 0 {
		putInt64Array(distMembers, array, macSid, pb_enum.PLATFORM_TYPE_MAC)
	}
	if windowsSid > 0 {
		putInt64Array(distMembers, array, windowsSid, pb_enum.PLATFORM_TYPE_WINDOWS)
	}
	if webSid > 0 {
		putInt64Array(distMembers, array, webSid, pb_enum.PLATFORM_TYPE_WEB)
	}
}

func putInt64Array(distMembers map[int64][]*pb_obj.Int64Array, array *pb_obj.Int64Array, sid int64, platform pb_enum.PLATFORM_TYPE) {
	m := &pb_obj.Int64Array{Vals: make([]int64, 4)}
	m.Vals[0] = sid // server_id
	m.Vals[1] = int64(platform)
	m.Vals[2] = array.GetUid()
	m.Vals[3] = array.GetStatus()
	distMembers[sid] = append(distMembers[sid], m)
}

func GetDistMembers(serverId int64, uid int64, status int64) (distMembers map[int64][]*pb_obj.Int64Array) {
	var (
		str   string
		array *pb_obj.Int64Array
	)
	distMembers = make(map[int64][]*pb_obj.Int64Array)
	str = fmt.Sprintf("%d,%d", serverId, status)
	array, serverId = pb_obj.MemberInt64Array(str, uid)
	setDistMembers(distMembers, serverId, array)
	return
}
