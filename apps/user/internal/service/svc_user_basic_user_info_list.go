package service

import (
	"context"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetBasicUserInfoList(ctx context.Context, req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp, err error) {
	resp = &pb_user.GetBasicUserInfoListResp{List: make([]*pb_user.BasicUserInfo, 0)}
	var (
		w = entity.NewMysqlWhere()
	)
	w.SetFilter("uid IN(?)", req.Uids)
	resp.List, err = s.userRepo.BasicUserInfoList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	xants.Submit(func() {
		s.cacheBasicUsers(resp.List)
	})
	return
}
