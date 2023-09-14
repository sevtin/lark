package svc_order

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_order"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_order"
	"lark/pkg/xhttp"
)

func (s *orderService) CreateRedEnvelopeOrder(params *dto_order.CreateRedEnvelopeOrderReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_order.CreateRedEnvelopeOrderReq)
		reply *pb_order.CreateRedEnvelopeOrderResp
	)
	_ = copier.Copy(req, params)
	req.Uid = uid
	reply = s.orderClient.CreateRedEnvelopeOrder(req)
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
	resp.Data = reply.PayUrl
	return
}
