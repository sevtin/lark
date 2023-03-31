package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
)

func (s *authService) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, _ error) {
	resp = new(pb_auth.SignOutResp)
	var (
		err error
	)
	err = s.userCache.SignOut(req.Uid, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_LOGOUT_FAILED, ERROR_AUTH_LOGOUT_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_LOGOUT_FAILED, ERROR_AUTH_LOGOUT_FAILED, err.Error())
		return
	}
	return
}
