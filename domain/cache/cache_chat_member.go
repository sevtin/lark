package cache

import (
	"context"
	"github.com/spf13/cast"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"strconv"
)

type ChatMemberCache interface {
	GetChatMemberInfo(chatId int64, uid int64) (info *pb_chat_member.ChatMemberInfo, err error)
	SetChatMemberInfo(info *pb_chat_member.ChatMemberInfo) (err error)

	HMSetChatMembers(chatId int64, remainder int, maps map[string]string) (err error)
	HSetNXChatMember(chatId int64, chatType pb_enum.CHAT_TYPE, uid int64, value string) (err error)

	HDelChatMembers(removes map[string][]string) (err error)
	HSetDistChatMembers(keys []string, field string, vals []string) (err error)
	GetDistChatMembers(chatId int64, remainder int) map[string]string
	GetChatMemberFlag(chatId int64, remainder int) (val string, err error)
	SetChatMemberFlag(chatId int64, remainder int, val string) (err error)
}

type chatMemberCache struct {
}

func NewChatMemberCache() ChatMemberCache {
	return &chatMemberCache{}
}

func (c *chatMemberCache) GetChatMemberInfo(chatId int64, uid int64) (info *pb_chat_member.ChatMemberInfo, err error) {
	var (
		key   = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(chatId)
		value string
	)
	info = new(pb_chat_member.ChatMemberInfo)
	value, err = xredis.HGet(key, cast.ToString(uid))
	if err != nil {
		return
	}
	if value != "" {
		err = utils.Unmarshal(value, info)
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
	err = xredis.CHMSet(key, map[string]string{cast.ToString(info.Uid): val}, constant.CONST_DURATION_CHAT_MEMBER_INFO_HASH_SECOND)
	return
}

func (c *chatMemberCache) GetChatMemberFlag(chatId int64, remainder int) (val string, err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_FLAG + utils.GetHashTagKey(chatId) + ":" + strconv.Itoa(remainder)
	)
	return xredis.Get(key)
}

func (c *chatMemberCache) SetChatMemberFlag(chatId int64, remainder int, val string) (err error) {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_FLAG + utils.GetHashTagKey(chatId) + ":" + strconv.Itoa(remainder)
	)
	return xredis.Set(key, val, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
}

func (c *chatMemberCache) HMSetChatMembers(chatId int64, remainder int, maps map[string]string) (err error) {
	if len(maps) == 0 {
		return
	}
	var (
		prefix  = xredis.GetPrefix()
		hashTag = utils.GetHashTagKey(chatId)
		slot    = ":" + strconv.Itoa(remainder)
		key1    = prefix + constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + hashTag + slot
		key2    = prefix + constant.RK_SYNC_DIST_CHAT_MEMBER_FLAG + hashTag + slot
		ctx     = context.Background()
	)
	pipe := xredis.Pipeline()
	pipe.HMSet(context.Background(), key1, maps)
	pipe.Expire(context.Background(), key1, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	pipe.Set(ctx, key2, constant.RV_SIGN_EXIST, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	_, err = pipe.Exec(ctx)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	return
}

func (c *chatMemberCache) HSetNXChatMember(chatId int64, chatType pb_enum.CHAT_TYPE, uid int64, value string) (err error) {
	var slot = ":0"
	if chatType != pb_enum.CHAT_TYPE_PRIVATE {
		slot = utils.GetChatSlot(uid)
	}
	var (
		prefix  = xredis.GetPrefix()
		hashTag = utils.GetHashTagKey(chatId)
		key1    = prefix + constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + hashTag + slot
		key2    = prefix + constant.RK_SYNC_DIST_CHAT_MEMBER_FLAG + hashTag + slot
		field   = cast.ToString(uid)
		ctx     = context.Background()
	)
	pipe := xredis.Pipeline()
	pipe.HSetNX(ctx, key1, field, value)
	pipe.Expire(ctx, key1, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	pipe.Set(ctx, key2, constant.RV_SIGN_EXIST, constant.CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND)
	_, err = pipe.Exec(ctx)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
		return
	}
	return
}

func (c *chatMemberCache) HDelChatMembers(removes map[string][]string) (err error) {
	err = xredis.CHDel(removes)
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
//	return xredis.HMGet(key, cast.ToString(uid))
//}

func (c *chatMemberCache) GetDistChatMembers(chatId int64, remainder int) map[string]string {
	var (
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chatId) + ":" + strconv.Itoa(remainder)
	)
	return xredis.HGetAll(key)
}
