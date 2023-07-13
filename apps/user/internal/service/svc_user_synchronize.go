package service

import (
	"gorm.io/gorm"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/protocol"
	"lark/pkg/utils"
)

func (s *userService) updateChatMemberCacheInfo(tx *gorm.DB, uid int64) (r *protocol.Result, err error) {
	r = new(protocol.Result)
	var (
		w       = entity.NewMysqlQuery()
		members []*pb_chat_member.ChatMemberInfo
		member  *pb_chat_member.ChatMemberInfo
		uidStr  = utils.Int64ToStr(uid)
		jsonStr string
		key     string
		keys    []string
		vals    []string
		index   int
	)
	w.SetFilter("uid=?", uid)
	w.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
	members, err = s.chatMemberRepo.TxChatMemberList(tx, w)
	if err != nil {
		r.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		return
	}

	keys = make([]string, len(members))
	vals = make([]string, len(members))
	for index, member = range members {
		jsonStr, err = utils.Marshal(member)
		if err != nil {
			r.Set(ERROR_CODE_USER_MARSHAL_FAILED, ERROR_USER_MARSHAL_FAILED)
			return
		}
		key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(member.ChatId)
		keys[index] = key
		vals[index] = jsonStr
	}
	if len(keys) == 0 {
		return
	}
	err = s.chatMemberCache.HSetDistChatMembers(keys, uidStr, vals)
	if err != nil {
		r.Set(ERROR_CODE_USER_CACHE_CHAT_MEMBER_INFO_FAILED, ERROR_USER_CACHE_CHAT_MEMBER_INFO_FAILED)
	}
	return
}
