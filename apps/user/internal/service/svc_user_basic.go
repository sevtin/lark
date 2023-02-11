package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetBasicUserInfo(ctx context.Context, req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp, _ error) {
	resp = &pb_user.GetBasicUserInfoResp{UserInfo: &pb_user.BasicUserInfo{}}
	var (
		w    = entity.NewMysqlWhere()
		user *pb_user.BasicUserInfo
		err  error
	)
	w.SetFilter("uid=?", req.Uid)
	user, err = s.userRepo.BasicUserInfo(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	xants.Submit(func() {
		s.userCache.SetBasicUserInfo(user)
	})
	return
}

func (s *userService) cacheBasicUsers(list []*pb_user.BasicUserInfo) {
	err := s.userCache.SetBasicUserInfoList(s.cfg.Redis.Prefix, list)
	if err != nil {
		xlog.Warn(ERROR_CODE_USER_REDIS_SET_FAILED, ERROR_USER_REDIS_SET_FAILED, err.Error())
	}
}
