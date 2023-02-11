package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (s *authService) SignUp(params *dto_auth.SignUpReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req      = new(pb_auth.SignUpReq)
		reply    *pb_auth.SignUpResp
		authResp = new(dto_auth.AuthResp)
		server   *dto_auth.ServerInfo
	)

	copier.Copy(req, params)
	server = s.getWsServer()
	req.ServerId = server.ServerId

	reply = s.authClient.SignUp(req)
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

	copier.Copy(authResp, reply)
	authResp.Server = server
	resp.Data = authResp
	return
}

func (s *authService) getWsServer() (wsServer *dto_auth.ServerInfo) {
	var (
		member   string
		server   string
		serverId int64
		port     int
	)
	list := s.serverMgrCache.ZRevRangeMsgGateway(0, 0)
	if len(list) > 0 {
		member = list[0]
	}
	server, serverId, port = utils.GetMsgGatewayServer(member)
	wsServer = &dto_auth.ServerInfo{
		ServerId: serverId,
		Name:     server,
		Port:     port,
	}
	return
}
