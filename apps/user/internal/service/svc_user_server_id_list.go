package service

import (
	"context"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetServerIdList(ctx context.Context, req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp, _ error) {
	resp = new(pb_user.GetServerIdListResp)
	var (
		w   = entity.NewMysqlWhere()
		err error
	)
	w.SetFilter("uid IN(?)", req.Uids)
	resp.List, err = s.userRepo.UserServerList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	xants.Submit(func() {
		s.userCache.SetUserServerList(resp.List)
	})
	return
}
