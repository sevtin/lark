package svc_chat_msg

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/xhttp"
)

func (s *chatMessageService) Search(req *dto_chat_msg.SearchMessageReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat_msg.SearchMessageReq)
		reply   *pb_chat_msg.SearchMessageResp
		res     = dto_chat_msg.SearchMessageResp{List: make([]*dto_chat_msg.MessageSummary, 0)}
	)
	copier.Copy(reqArgs, req)
	reqArgs.Uid = uid
	reply = s.chatMessageClient.SearchMessage(reqArgs)
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
	res.Total = reply.Total
	copier.Copy(&res.List, reply.List)
	resp.Data = res
	return
}
