package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func (s *userService) GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, _ error) {
	resp = &pb_user.GetUserListResp{List: make([]*pb_user.UserInfo, 0)}
	var (
		w    = entity.NewMysqlWhere()
		list []*po.User
		err  error
	)
	w.SetFilter("uid IN(?)", req.UidList)
	list, err = s.userRepo.UserList(w)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.List, list)
	return
}
