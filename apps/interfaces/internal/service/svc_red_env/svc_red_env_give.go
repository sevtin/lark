package svc_red_env

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_red_env"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_red_env"
	"lark/pkg/xhttp"
)

func (s *redEnvService) GiveRedEnvelope(params *dto_red_env.GiveRedEnvelopeReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_red_env.GiveRedEnvelopeReq)
		reply *pb_red_env.GiveRedEnvelopeResp
	)
	copier.Copy(req, params)
	req.SenderUid = uid
	reply = s.redEnvClient.GiveRedEnvelope(req)
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
