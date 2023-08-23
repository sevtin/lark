package svc_red_env_receive

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_red_env_receive"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_red_env_receive"
	"lark/pkg/xhttp"
)

func (s *redEnvReceiveService) OpenRedEnvelope(params *dto_red_env_receive.OpenRedEnvelopeReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_red_env_receive.OpenRedEnvelopeReq)
		reply *pb_red_env_receive.OpenRedEnvelopeResp
	)
	copier.Copy(req, params)
	req.Uid = uid
	reply = s.redEnvReceiveClient.OpenRedEnvelope(req)
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
	resp.Data = reply.Result
	return
}
