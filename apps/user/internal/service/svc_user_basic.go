package service

import (
	"context"
	"lark/domain/pdo"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetBasicUserInfo(ctx context.Context, req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp, _ error) {
	resp = &pb_user.GetBasicUserInfoResp{UserInfo: &pb_user.BasicUserInfo{}}
	var (
		user = new(pdo.BasicUserInfo)
		q    = entity.NewMysqlQuery()
		err  error
	)
	q.Fields = user.GetField()
	q.SetFilter("uid=?", req.Uid)
	err = s.userRepo.QueryUser(q, user)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	if resp.UserInfo.Uid == 0 {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		return
	}
	resp.UserInfo = &pb_user.BasicUserInfo{
		Uid:      user.Uid,
		LarkId:   user.LarkId,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		BirthTs:  user.BirthTs,
		CityId:   user.CityId,
		Avatar:   user.Avatar,
	}
	xants.Submit(func() {
		s.userCache.SetBasicUserInfo(resp.UserInfo)
	})
	return
}

func (s *userService) cacheBasicUsers(list []*pb_user.BasicUserInfo) {
	err := s.userCache.SetBasicUserInfoList(list)
	if err != nil {
		xlog.Warn(ERROR_CODE_USER_REDIS_SET_FAILED, ERROR_USER_REDIS_SET_FAILED, err.Error())
	}
}
