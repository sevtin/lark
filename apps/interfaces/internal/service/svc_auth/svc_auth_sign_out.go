package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

func (s *authService) SignOut(params *dto_auth.SignOutReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_auth.SignOutReq)
		reply *pb_auth.SignOutResp
	)
	copier.Copy(req, params)
	reply = s.authClient.SignOut(req)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
	}
	return
}
