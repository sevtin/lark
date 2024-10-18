package logic

import (
	"github.com/spf13/cast"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"sync"
)

func GetMembersFromHash(hashmap map[string]string) (distMembers map[int64][]*pb_obj.Int64Array) {
	if len(hashmap) == 0 {
		return
	}
	return groupFromHashmap(hashmap)
}

func groupFromHashmap(hashmap map[string]string) (distMembers map[int64][]*pb_obj.Int64Array) {
	var (
		srvIds    = getUsersServerId(hashmap)
		uid       string
		uidVal    int64
		status    string
		statusVal int64
		array     *pb_obj.Int64Array
		serverId  int64
	)
	distMembers = make(map[int64][]*pb_obj.Int64Array)
	for uid, status = range hashmap {
		uidVal = cast.ToInt64(uid)
		serverId, _ = srvIds[uidVal]
		statusVal = cast.ToInt64(status)
		array = pb_obj.MemberInt64Array(uidVal, statusVal, serverId)
		setDistMembers(distMembers, serverId, array)
	}
	return
}

func getUsersServerId(hashmap map[string]string) (srvIds map[int64]int64) {
	var (
		key   = constant.RK_SYNC_USER_SERVER
		uid   string
		slots = map[int64][]string{}
		slot  int64
		uids  []string
	)
	srvIds = make(map[int64]int64)
	for uid, _ = range hashmap {
		slot = utils.UserSlot(cast.ToInt64(uid))
		slots[slot] = append(slots[slot], uid)
	}
	wg := &sync.WaitGroup{}
	lock := &sync.RWMutex{}
	for slot, uids = range slots {
		wg.Add(1)
		go func(w *sync.WaitGroup, s int64, ids []string) {
			defer w.Done()
			servers := xredis.HMGet(key+cast.ToString(s), ids...)
			lock.Lock()
			for i, sid := range servers {
				srvIds[cast.ToInt64(ids[i])] = cast.ToInt64(sid)
			}
			lock.Unlock()
		}(wg, slot, uids)
	}
	wg.Wait()
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
		array *pb_obj.Int64Array
	)
	distMembers = make(map[int64][]*pb_obj.Int64Array)
	array = pb_obj.MemberInt64Array(uid, serverId, status)
	setDistMembers(distMembers, serverId, array)
	return
}
