package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func (s *chatMemberService) GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, _ error) {
	resp = &pb_chat_member.GetChatMemberListResp{}
	var (
		w   = entity.NewMysqlWhere()
		err error
	)
	w.SetFilter("chat_id=?", req.ChatId)
	w.SetFilter("uid>?", req.LastUid)
	w.SetSort("role_id DESC,uid ASC")
	w.SetLimit(req.Limit)
	resp.List, err = s.chatMemberRepo.GroupChatMemberInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
