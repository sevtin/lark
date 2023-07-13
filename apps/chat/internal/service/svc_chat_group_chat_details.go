package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/cr/cr_user"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
)

func (s *chatService) GroupChatDetails(ctx context.Context, req *pb_chat.GroupChatDetailsReq) (resp *pb_chat.GroupChatDetailsResp, _ error) {
	resp = &pb_chat.GroupChatDetailsResp{Details: &pb_chat.GroupChatDetails{Creator: &pb_chat.ChatCreator{}}}
	var (
		w    = entity.NewMysqlQuery()
		chat *po.Chat
		user *pb_user.BasicUserInfo
		err  error
	)
	w.SetFilter("chat_id=?", req.ChatId)
	w.SetFilter("chat_type=?", pb_enum.CHAT_TYPE_GROUP)
	chat, err = s.chatRepo.Chat(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}
	if chat.ChatId == 0 {
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, req.ChatId)
		return
	}
	user, err = cr_user.GetBasicUserInfo(s.userCache, s.userRepo, chat.CreatorUid)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, chat.CreatorUid)
		return
	}
	copier.Copy(resp.Details, chat)
	copier.Copy(resp.Details.Creator, user)
	return
}
