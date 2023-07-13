package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/do"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
)

func (s *chatInviteService) ChatInviteList(_ context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, _ error) {
	resp = &pb_invite.ChatInviteListResp{List: make([]*pb_invite.ChatInviteInfo, 0)}
	var (
		w    = entity.NewMysqlQuery()
		list []*do.ChatInvite
		err  error
	)
	w.Limit = int(req.Limit)
	w.SetSort("invite_id DESC")
	if req.MaxInviteId > 0 {
		w.SetFilter("invite_id<?", req.MaxInviteId)
	}
	if req.HandleResult > 0 {
		w.SetFilter("handle_result=?", req.HandleResult)
	}
	switch req.Role {
	case pb_enum.INVITE_ROLE_INITIATOR: // 发起者
		w.SetFilter("initiator_uid=?", req.Uid)
	case pb_enum.INVITE_ROLE_APPROVER: // 审批人
		w.SetFilter("invitee_uid=?", req.Uid)
	}
	list, err = s.chatInviteRepo.ChatInviteList(w)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_INVITE_QUERY_DB_FAILED, ERROR_CHAT_INVITE_QUERY_DB_FAILED, err)
		return
	}
	copier.Copy(&resp.List, list)
	return
}
