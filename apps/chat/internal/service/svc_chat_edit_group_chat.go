package service

import (
	"context"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
)

func (s *chatService) EditGroupChat(ctx context.Context, req *pb_chat.EditGroupChatReq) (resp *pb_chat.EditGroupChatResp, _ error) {
	resp = &pb_chat.EditGroupChatResp{}
	var (
		w      = entity.NewMysqlWhere()
		u      = entity.NewMysqlUpdate()
		member *pb_chat_member.ChatMemberInfo
		chat   *po.Chat
		err    error
	)
	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()
	w.SetFilter("chat_id=?", req.ChatId)
	w.SetFilter("uid=?", req.Uid)
	w.SetFilter("role_id>=?", int32(pb_enum.CHAT_GROUP_ROLE_ADMINISTRATOR))
	member, err = s.chatMemberRepo.ChatMember(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		//xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}
	if member.ChatId == 0 {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		//xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		return
	}
	if member.RoleId == 0 {
		resp.Set(ERROR_CODE_CHAT_NO_RIGHT_TO_MODIFY, ERROR_CHAT_NO_RIGHT_TO_MODIFY)
		//xlog.Warn(ERROR_CODE_CHAT_NO_RIGHT_TO_MODIFY, ERROR_CHAT_NO_RIGHT_TO_MODIFY)
		return
	}

	w.Reset()
	w.SetFilter("chat_id=?", req.ChatId)
	chat, err = s.chatRepo.Chat(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		//xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}

	u.SetFilter("chat_id=?", req.ChatId)
	req.Kvs.StrFieldValidation(chatUpdateFields, u.Values)
	var (
		name     interface{}
		about    interface{}
		ok       bool
		isUpdate bool
	)
	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		err = s.chatRepo.TxUpdateChat(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
			//xlog.Warn(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED, err.Error())
			return
		}

		if name, ok = u.Values["name"]; ok {
			if name.(string) != chat.Name {
				isUpdate = true
				u.Reset()
				u.SetFilter("chat_id=?", req.ChatId)
				u.Set("chat_name", name)
				err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
				if err != nil {
					resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
					//xlog.Warn(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED, err.Error())
					return
				}
			}
		}
		return
	})
	if err != nil {
		return
	}
	if about, ok = u.Values["about"]; ok {
		if about.(string) != chat.About {
			isUpdate = true
		}
	}
	if isUpdate == true && err == nil {
		err = s.chatCache.DelChatInfo(chat.ChatId)
	}
	return
}
