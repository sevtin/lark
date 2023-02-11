package svc_chat_member

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_member"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/xhttp"
)

func (s *chatMemberService) ChatMemberList(params *dto_chat_member.ChatMemberListReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = &pb_chat_member.GetChatMemberListReq{}
		reply *pb_chat_member.GetChatMemberListResp
	)
	copier.Copy(req, params)

	reply = s.chatMemberClient.GetChatMemberList(req)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		return
	}
	resp.Data = reply.List
	return
}
