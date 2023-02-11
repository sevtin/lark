package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/xhttp"
)

func (s *authService) SignIn(params *dto_auth.SignInReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req       = new(pb_auth.SignInReq)
		reply     *pb_auth.SignInResp
		loginResp = new(dto_auth.AuthResp)
		server    *dto_auth.ServerInfo
		onOffResp *pb_chat_member.ChatMemberOnOffLineResp
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
	server = s.getWsServer()
	onOffResp = s.chatMemberOnOffLine(reply.UserInfo.Uid, server.ServerId, req.Platform)
	if onOffResp == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if onOffResp.Code > 0 {
		resp.SetResult(onOffResp.Code, onOffResp.Msg)
		xlog.Warn(onOffResp.Code, onOffResp.Msg)
		return
	}
	copier.Copy(loginResp, reply)
	loginResp.Server = server
	resp.Data = loginResp
	return
}

func (s *authService) chatMemberOnOffLine(uid int64, serverId int64, platform pb_enum.PLATFORM_TYPE) (resp *pb_chat_member.ChatMemberOnOffLineResp) {
	req := &pb_chat_member.ChatMemberOnOffLineReq{
		Uid:      uid,
		ServerId: serverId,
		Platform: platform,
	}
	return s.chatMemberClient.ChatMemberOnOffLine(req)
}
