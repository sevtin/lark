package svc_chat_member

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_member"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/xhttp"
)

func (s *chatMemberService) GroupChatList(params *dto_chat_member.GroupChatListReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = &pb_chat_member.GetGroupChatListReq{}
		reply   *pb_chat_member.GetGroupChatListResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid

	reply = s.chatMemberClient.GetGroupChatList(reqArgs)
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
