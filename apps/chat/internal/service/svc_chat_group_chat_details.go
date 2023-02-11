package service

import (
	"context"
	"github.com/jinzhu/copier"
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
		w       = entity.NewMysqlWhere()
		chat    *po.Chat
		reqArgs *pb_user.GetBasicUserInfoReq
		reply   *pb_user.GetBasicUserInfoResp
		err     error
	)
	w.SetFilter("chat_id=?", req.ChatId)
	w.SetFilter("chat_type=?", int32(pb_enum.CHAT_TYPE_GROUP))
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
	reqArgs = &pb_user.GetBasicUserInfoReq{Uid: chat.CreatorUid}
	reply = s.userClient.GetBasicUserInfo(reqArgs)
	if reply == nil {
		resp.Set(ERROR_CODE_CHAT_GRPC_SERVICE_FAILURE, ERROR_CHAT_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_CHAT_GRPC_SERVICE_FAILURE, ERROR_CHAT_GRPC_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.Set(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	if reply.UserInfo.Uid == 0 {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, chat.CreatorUid)
	}
	copier.Copy(resp.Details, chat)
	copier.Copy(resp.Details.Creator, reply.UserInfo)
	return
}
