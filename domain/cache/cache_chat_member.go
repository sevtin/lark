package cache

import (
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

type ChatMemberCache interface {
	GetChatMemberInfo(chatId int64, uid int64) (info *pb_chat_member.ChatMemberInfo, err error)
	SetChatMemberInfo(info *pb_chat_member.ChatMemberInfo) (err error)

	HMSetChatMembers(chatId int64, maps map[string]string) (err error)
	HSetNXChatMember(chatId int64, uid int64, value string) (err error)

	HDelChatMembers(keys []string, fields []string) (err error)
	HSetDistChatMembers(keys []string, field string, vals []string) (err error)
	//GetDistChatMember(chatId int64, uid int64) []interface{}
	GetAllDistChatMembers(chatId int64) map[string]string
}

type chatMemberCache struct {
}

func NewChatMemberCache() ChatMemberCache {
	return &chatMemberCache{}
}

func (c *chatMemberCache) GetChatMemberInfo(chatId int64, uid int64) (info *pb_chat_member.ChatMemberInfo, err error) {
	var (
		key    = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(chatId)
		values []interface{}
	)
	info = new(pb_chat_member.ChatMemberInfo)
	values = xredis.HMGet(key, utils.Int64ToStr(uid))
	if len(values) == 1 && values[0] != nil {
		switch values[0].(type) {
		case string:
			err = utils.Unmarshal(values[0].(string), info)
		}
	}
	return
}

func (c *chatMemberCache) SetChatMemberInfo(info *pb_chat_member.ChatMemberInfo) (err error) {
	var (
		key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(info.ChatId)
		val string
	)
	val, err = utils.Marshal(info)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	err = xredis.CHMSet(key, map[string]interface{}{utils.Int64ToStr(info.Uid): val}, constant.CONST_DURATION_CHAT_MEMBER_INFO_HASH_SECOND)
	return
}

func (c *chatMemberCache) HMSetChatMembers(chatId int64, maps map[string]string) (err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chatId)
	)
	err = xredis.HMSet(key, maps)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	return
}

func (c *chatMemberCache) HSetNXChatMember(chatId int64, uid int64, value string) (err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chatId)
	)
	err = xredis.HSetNX(key, utils.Int64ToStr(uid), value)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	return
}

func (c *chatMemberCache) HDelChatMembers(keys []string, fields []string) (err error) {
	err = xredis.CHDel(keys, fields)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_DELETE_FAILED, ERROR_CACHE_REDIS_DELETE_FAILED, err.Error())
	}
	return
}

func (c *chatMemberCache) HSetDistChatMembers(keys []string, field string, vals []string) (err error) {
	err = xredis.CBatchHSet(keys, field, vals)
	return
}

//func (c *chatMemberCache) GetDistChatMember(chatId int64, uid int64) []interface{} {
//	var (
//		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chatId)
//	)
//	return xredis.HMGet(key, utils.Int64ToStr(uid))
//}

func (c *chatMemberCache) GetAllDistChatMembers(chatId int64) map[string]string {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chatId)
	)
	return xredis.HGetAll(key)
}
