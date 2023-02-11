package cache

import (
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/utils"
)

type ChatCache interface {
	GetGroupChatInfo(chatId int64) (info *pb_chat.ChatInfo, err error)
	SetGroupChatInfo(info *pb_chat.ChatInfo) (err error)
	DelChatInfo(chatId int64) (err error)
	// GetGroupChatDetails(chatId int64) (details *pb_chat.GroupChatDetails, err error)
	// SetGroupChatDetails(details *pb_chat.GroupChatDetails) (err error)
}

type chatCache struct {
}

func NewChatCache() ChatCache {
	return &chatCache{}
}

func (c *chatCache) GetGroupChatInfo(chatId int64) (info *pb_chat.ChatInfo, err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.Int64ToStr(chatId)
	)
	info = new(pb_chat.ChatInfo)
	err = Get(key, info)
	return
}

func (c *chatCache) SetGroupChatInfo(info *pb_chat.ChatInfo) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.Int64ToStr(info.ChatId)
	)
	err = Set(key, info, constant.CONST_DURATION_GROUP_CHAT_INFO_SECOND)
	return
}

func (c *chatCache) DelGroupChatInfo(chatId int64) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.Int64ToStr(chatId)
	)
	err = xredis.Del(key)
	return
}

func (c *chatCache) DelChatInfo(chatId int64) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.Int64ToStr(chatId)
	)
	err = xredis.Del(key)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_DELETE_FAILED, ERROR_CACHE_REDIS_DELETE_FAILED, err.Error())
	}
	return
}

//func (c *chatCache) GetGroupChatDetails(chatId int64) (details *pb_chat.GroupChatDetails, err error) {
//	var (
//		key = constant.RK_SYNC_GROUP_CHAT_DETAILS + utils.Int64ToStr(chatId)
//	)
//	details = new(pb_chat.GroupChatDetails)
//	err = Get(key, details)
//	return
//}
//
//func (c *chatCache) SetGroupChatDetails(details *pb_chat.GroupChatDetails) (err error) {
//	var (
//		key = constant.RK_SYNC_GROUP_CHAT_DETAILS + utils.Int64ToStr(details.ChatId)
//	)
//	err = Set(key, details, constant.CONST_DURATION_GROUP_CHAT_INFO_SECOND)
//	return
//}
