package service

import (
	"gorm.io/gorm"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
)

func (s *chatService) removeChatMember(u *entity.MysqlUpdate, chatId int64, uidList []int64, chatType pb_enum.CHAT_TYPE) (rowsAffected int64, err error) {
	var (
		key1     = s.cfg.Redis.Prefix + constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
		key2     = s.cfg.Redis.Prefix + constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.Int64ToStr(chatId)
		uidCount = len(uidList)
		keys     []string
		fields   []interface{}
		uid      int64
		index    int
	)
	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		rowsAffected, err = s.chatMemberRepo.TxQuitChatMember(tx, u)
		if err != nil {
			return
		}
		switch chatType {
		case pb_enum.CHAT_TYPE_PRIVATE:
			if rowsAffected != 1 {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
			if uidCount != 2 {
				return
			}
			u.Reset()
			u.SetFilter("chat_id=?", chatId)
			u.SetFilter("owner_id=?", uidList[1])
			u.SetFilter("deleted_ts=?", 0)
			u.Set("status", int32(pb_enum.CHAT_STATUS_NON_CONTACT))
			rowsAffected, err = s.chatMemberRepo.TxQuitChatMember(tx, u)
			if err != nil {
				return
			}
			if rowsAffected != 1 {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
		case pb_enum.CHAT_TYPE_GROUP:
			if int(rowsAffected) != len(uidList) {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
		}
		return
	})

	keys = make([]string, uidCount*2)
	fields = make([]interface{}, uidCount*2)
	for index, uid = range uidList {
		keys[index] = key1
		fields[index] = uid

		keys[index+uidCount] = key2
		fields[index+uidCount] = uid
	}
	if len(keys) == 0 {
		return
	}
	err = s.chatMemberCache.HDelChatMembers(keys, fields)
	return
}
