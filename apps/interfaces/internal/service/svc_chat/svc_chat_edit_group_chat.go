package svc_chat

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/xhttp"
)

func (s *chatService) EditGroupChat(params *dto_chat.EditGroupChatReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat.EditGroupChatReq)
		reply   *pb_chat.EditGroupChatResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid
	reply = s.chatClient.EditGroupChat(reqArgs)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	return
}
