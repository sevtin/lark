package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

func (s *authService) GoogleOAuth2Callback(params *dto_auth.GoogleOauthCallbackReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_auth.GoogleOAuth2CallbackReq)
		reply *pb_auth.GoogleOAuth2CallbackResp
	)
	_ = copier.Copy(req, params)
	reply = s.authClient.GoogleOAuth2Callback(req)
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
	resp.Data = reply.AuthUserInfo
	return
}
