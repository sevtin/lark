package svc_convo

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_convo"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/xhttp"
)

func (s *convoService) ConvoList(req *dto_convo.ConvoListReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_convo.ConvoListReq)
		reply   *pb_convo.ConvoListResp
	)
	copier.Copy(reqArgs, req)
	reply = s.convoClient.ConvoList(reqArgs)
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
	resp.Data = reply.List
	return
}
