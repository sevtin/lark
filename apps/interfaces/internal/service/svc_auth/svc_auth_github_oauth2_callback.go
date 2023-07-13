package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

func (s *authService) GithubOAuth2Callback(params *dto_auth.GithubOauthCallbackReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req       = new(pb_auth.GithubOAuth2CallbackReq)
		reply     *pb_auth.GithubOAuth2CallbackResp
		loginResp = new(dto_auth.AuthResp)
	)
	copier.Copy(req, params)
	reply = s.authClient.GithubOAuth2Callback(req)
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
