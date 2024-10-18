package utils

import (
	"lark/pkg/constant"
	"strconv"
)

func GetChatSlot(uid int64) string {
	return ":" + strconv.FormatInt(uid%constant.MAX_CHAT_SLOT, 10)
}

func ChatSlot(uid int64) int64 {
	return uid % constant.MAX_CHAT_SLOT
}

func GetUserSlot(uid int64) string {
	return strconv.FormatInt(uid%constant.MAX_USER_SLOT, 10)
}

func UserSlot(uid int64) int64 {
	return uid % constant.MAX_USER_SLOT
}
