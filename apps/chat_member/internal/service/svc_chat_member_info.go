package service

import (
	"context"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

func (s *chatMemberService) GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, _ error) {
	resp = &pb_chat_member.GetChatMemberInfoResp{Info: new(pb_chat_member.ChatMemberInfo)}
	var (
		w   = entity.NewNormalWhere()
		err error
	)
	w.SetFilter("chat_id = ?", req.ChatId)
	w.SetFilter("uid = ?", req.Uid)
	resp.Info, err = s.chatMemberRepo.ChatMember(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	if resp.Info.Uid == 0 {
		resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
		return
	}
	xants.Submit(func() {
		s.chatMemberCache.SetChatMemberInfo(resp.Info)
	})
	return
}
