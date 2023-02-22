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

	HMSetChatMembers(chatId int64, maps map[string]interface{}) (err error)
	HSetNXChatMember(chatId int64, uid int64, value string) (err error)

	HDelChatMembers(keys []string, fields []interface{}) (err error)
	HMSetDistChatMembers(keys []string, vals []interface{}) (err error)
	GetDistChatMember(chatId int64, uid int64) []interface{}
	GetAllDistChatMembers(chatId int64) map[string]string
}

type chatMemberCache struct {
}

func NewChatMemberCache() ChatMemberCache {
	return &chatMemberCache{}
}

func (c *chatMemberCache) GetChatMemberInfo(chatId int64, uid int64) (info *pb_chat_member.ChatMemberInfo, err error) {
	var (
		key  = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.Int64ToStr(chatId)
		vals []interface{}
	)
	info = new(pb_chat_member.ChatMemberInfo)
	vals = xredis.HMGet(key, utils.Int64ToStr(uid))
	if len(vals) == 1 && vals[0] != nil {
		err = utils.Unmarshal(vals[0].(string), info)
	}
	return
}

func (c *chatMemberCache) SetChatMemberInfo(info *pb_chat_member.ChatMemberInfo) (err error) {
	var (
		key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.Int64ToStr(info.ChatId)
	)
	err = HMSet(key, utils.Int64ToStr(info.Uid), info)
	if err != nil {
		return
	}
	err = xredis.Expire(key, constant.CONST_DURATION_CHAT_MEMBER_INFO_HASH_SECOND)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_SET_EXPIRE_FAILED, ERROR_CACHE_SET_EXPIRE_FAILED, err.Error())
		return
	}
	return
}

func (c *chatMemberCache) HMSetChatMembers(chatId int64, maps map[string]interface{}) (err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
	)
	err = xredis.HMSet(key, maps)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	//err = xredis.Expire(key, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	//if err != nil {
	//	xlog.Warn(ERROR_CODE_CACHE_SET_EXPIRE_FAILED, ERROR_CACHE_SET_EXPIRE_FAILED, err.Error())
	//}
	return
}

func (c *chatMemberCache) HSetNXChatMember(chatId int64, uid int64, value string) (err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
	)
	err = xredis.HSetNX(key, utils.Int64ToStr(uid), value)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	//err = xredis.Expire(key, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	//if err != nil {
	//	xlog.Warn(ERROR_CODE_CACHE_SET_EXPIRE_FAILED, ERROR_CACHE_SET_EXPIRE_FAILED, err.Error())
	//	return
	//}
	return
}

func (c *chatMemberCache) HDelChatMembers(keys []string, fields []interface{}) (err error) {
	if len(keys) == 0 || len(fields) == 0 {
		return
	}
	err = xredis.EvalSha(xredis.SHA_HDEL_CHAT_MEMBER, keys, fields)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_DELETE_FAILED, ERROR_CACHE_REDIS_DELETE_FAILED, err.Error())
	}
	return
}

func (c *chatMemberCache) HMSetDistChatMembers(keys []string, vals []interface{}) (err error) {
	err = xredis.EvalSha(xredis.SHA_HMSET_DIST_CHAT_MEMBER, keys, vals)
	return
}

func (c *chatMemberCache) GetDistChatMember(chatId int64, uid int64) []interface{} {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
	)
	return xredis.HMGet(key, utils.Int64ToStr(uid))
}

func (c *chatMemberCache) GetAllDistChatMembers(chatId int64) map[string]string {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
	)
	return xredis.HGetAll(key)
}
