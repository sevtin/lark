package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
)

func (s *chatMemberService) GetGroupChatList(ctx context.Context, req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp, _ error) {
	resp = new(pb_chat_member.GetGroupChatListResp)
	var (
		w   = entity.NewMysqlWhere()
		err error
	)
	w.SetFilter("uid=?", req.Uid)
	w.SetFilter("chat_type=?", pb_enum.CHAT_TYPE_GROUP)
	w.SetFilter("chat_id>?", req.LastChatId)
	w.SetLimit(req.Limit)
	resp.List, err = s.chatMemberRepo.GroupChatBasicInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
