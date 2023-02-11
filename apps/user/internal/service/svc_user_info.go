package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	resp = &pb_user.UserInfoResp{}
	var (
		err error
	)
	resp.UserInfo, _ = s.userCache.GetUserInfo(req.Uid)
	if resp.UserInfo.Uid > 0 {
		return
	}
	err = s.queryUserInfo(req.Uid, resp)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	if resp.UserInfo.Uid == 0 {
		return
	}
	err = s.queryUserAvatar(resp.UserInfo.Uid, resp)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	xants.Submit(func() {
		s.userCache.SetUserInfo(resp.UserInfo)
	})
	return
}

func (s *userService) queryUserInfo(uid int64, resp *pb_user.UserInfoResp) (err error) {
	var (
		w    = entity.NewMysqlWhere()
		user *po.User
	)
	w.SetFilter("uid = ?", uid)
	user, err = s.userRepo.UserInfo(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	if user.Uid == 0 {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, user.Uid)
		return
	}
	copier.Copy(resp.UserInfo, user)
	return
}

func (s *userService) queryUserAvatar(uid int64, resp *pb_user.UserInfoResp) (err error) {
	var (
		w      = entity.NewMysqlWhere()
		avatar *po.Avatar
	)
	w.SetFilter("owner_id = ?", uid)
	w.SetFilter("owner_type = ?", int32(pb_enum.CHAT_TYPE_PRIVATE))
	avatar, err = s.avatarRepo.Avatar(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo.Avatar, avatar)
	return
}
