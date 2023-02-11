package utils

import (
	"lark/pkg/proto/pb_enum"
	"strings"
)

func MemberHash(uid1 int64, uid2 int64) string {
	var (
		sessionKey string
	)
	if uid1 > uid2 {
		sessionKey = Int64ToStr(uid1) + "-" + Int64ToStr(uid2)
	} else {
		sessionKey = Int64ToStr(uid2) + "-" + Int64ToStr(uid1)
	}
	return MD5(sessionKey)
}

func ChatStatus(in interface{}) (status pb_enum.CHAT_STATUS) {
	var (
		str string
		arr []string
		val int
	)
	str = ToString(in)
	arr = strings.Split(str, ",")
	if len(arr) == 4 {
		val, _ = ToInt(arr[3])
		status = pb_enum.CHAT_STATUS(val)
	}
	return
}

func StrUidList(uidList []int64) (list []string) {
	var (
		i   int
		uid int64
	)
	list = make([]string, len(uidList))
	for i, uid = range uidList {
		list[i] = Int64ToStr(uid)
	}
	return
}
