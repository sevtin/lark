package service

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
	"lark/pkg/protocol"
	"strings"
)

func (s *userService) EditUserInfo(ctx context.Context, req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp, _ error) {
	resp = new(pb_user.EditUserInfoResp)
	var (
		u        = entity.NewMysqlUpdate()
		nickname interface{}
		ok       bool
		result   *protocol.Result
		err      error
	)
	u.SetFilter("uid=?", req.Uid)
	req.Kvs.StrFieldValidation(userUpdateFields, u.Values)
	req.Kvs.IntFieldValidation(userUpdateFields, u.Values)
	tx := xmysql.GetTX()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = s.userRepo.TxUpdateUser(tx, u)
	if err != nil {
		switch err.(type) {
		case *mysql.MySQLError:
			if err.(*mysql.MySQLError).Number == constant.ERROR_CODE_MYSQL_DUPLICATE_ENTRY {
				if strings.HasSuffix(err.(*mysql.MySQLError).Message, constant.DUPLICATE_ENTRY_KV_USERS_MOBILE) {
					resp.Set(ERROR_CODE_USER_MOBILE_HAS_BEEN_OCCUPIED, ERROR_USER_MOBILE_HAS_BEEN_OCCUPIED)
					xlog.Warn(ERROR_CODE_USER_MOBILE_HAS_BEEN_OCCUPIED, ERROR_USER_MOBILE_HAS_BEEN_OCCUPIED, err.Error())
					return
				}
				if strings.HasSuffix(err.(*mysql.MySQLError).Message, constant.DUPLICATE_ENTRY_KV_USERS_LARK_ID) {
					resp.Set(ERROR_CODE_USER_LARK_ID_HAS_BEEN_OCCUPIED, ERROR_USER_LARK_ID_HAS_BEEN_OCCUPIED)
					xlog.Warn(ERROR_CODE_USER_LARK_ID_HAS_BEEN_OCCUPIED, ERROR_USER_LARK_ID_HAS_BEEN_OCCUPIED, err.Error())
					return
				}
			}
		}
		resp.Set(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED, err.Error())
		return
	}

	// 删除缓存
	err = s.userCache.DelUserInfo(s.cfg.Redis.Prefix, req.Uid)
	if err != nil {
		resp.Set(ERROR_CODE_USER_UPDATE_USER_CACHE_FAILED, ERROR_USER_UPDATE_USER_CACHE_FAILED)
		xlog.Warn(ERROR_CODE_USER_UPDATE_USER_CACHE_FAILED, ERROR_USER_UPDATE_USER_CACHE_FAILED, err.Error())
		return
	}
	if nickname, ok = u.Values["nickname"]; ok == false {
		return
	}
	u.Reset()
	u.SetFilter("uid=?", req.Uid)
	u.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
	u.Set("alias", nickname)
	err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
	if err != nil {
		resp.Set(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	result, err = s.updateChatMemberCacheInfo(tx, req.Uid)
	if err != nil {
		resp.Set(result.Code, result.Msg)
		return
	}
	return
}
