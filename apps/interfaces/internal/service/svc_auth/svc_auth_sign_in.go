package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

func (s *authService) SignIn(params *dto_auth.SignInReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req       = new(pb_auth.SignInReq)
		reply     *pb_auth.SignInResp
		loginResp = new(dto_auth.AuthResp)
	)
	copier.Copy(req, params)
	reply = s.authClient.SignIn(req)
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
	copier.Copy(loginResp, reply)
	resp.Data = loginResp
	return
}
