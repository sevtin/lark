package svc_user

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_user"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_user"
	"lark/pkg/xhttp"
)

func (s *userService) GetUserInfo(params *dto_user.UserInfoReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_user.UserInfoReq)
		reply *pb_user.UserInfoResp
	)
	copier.Copy(req, params)
	if params.IsSelf == true {
		req.Uid = uid
	} else {
		req.Uid = params.Uid
	}
	reply = s.userClient.GetUserInfo(req)
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
	resp.Data = reply.UserInfo
	return
}
