package service

import (
	"context"
	"fmt"
	"lark/domain/do"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *chatMemberService) ChatMemberOnOffLine(ctx context.Context, req *pb_chat_member.ChatMemberOnOffLineReq) (resp *pb_chat_member.ChatMemberOnOffLineResp, _ error) {
	resp = new(pb_chat_member.ChatMemberOnOffLineResp)
	var (
		err error
	)
	err = s.memberOnOffLine(req.Uid, req.ServerId, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED, ERROR_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED, ERROR_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED, err.Error())
		return
	}
	return
}

func (s *chatMemberService) memberOnOffLine(uid int64, serverId int64, platform pb_enum.PLATFORM_TYPE) (err error) {
	var (
		oldSidStr string
		oldSidVal int64
	)
	// 1、获取旧的serverId
	oldSidStr, err = s.userCache.GetServerId(uid)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_REDIS_GET_FAILED, ERROR_CHAT_MEMBER_REDIS_GET_FAILED, err.Error())
	}
	if oldSidStr == "" {
		var (
			user *pb_user.UserServerId
			w    = entity.NewMysqlWhere()
		)
		w.SetFilter("uid=?", uid)
		user, err = s.userRepo.UserServerId(w)
		if err != nil {
			xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
			return
		}
		oldSidVal = user.ServerId
	} else {
		oldSidVal = utils.StrToInt64(oldSidStr)
	}
	// 2、得到新的serverId
	serverId = utils.NewServerId(oldSidVal, serverId, platform)
	// 3、是否是同一台服务器
	if oldSidVal == serverId {
		return
	}
	// 4、更新数据库和ChatMember缓存
	err = s.updateMemberConnectedServer(uid, serverId)
	if err != nil {
		return
	}
	// 5、更新serverId缓存
	err = s.userCache.SetServerId(uid, serverId)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_REDIS_SET_FAILED, ERROR_CHAT_MEMBER_REDIS_SET_FAILED, err.Error())
		err = nil
	}
	return
}

func (s *chatMemberService) updateMemberConnectedServer(uid int64, serverId int64) (err error) {
	var (
		u         = entity.NewMysqlUpdate()
		w         = entity.NewMysqlWhere()
		allStatus []*do.ChatMemberStatus
		list      []*do.ChatMemberStatus
		limit     = 2000
		maxChatId int64
	)

	defer func() {
		if err != nil {
			xlog.Warn(err.Error())
		}
	}()

	// 1 获取Chat成员
	for {
		w.Reset()
		w.SetFilter("uid = ?", uid)
		w.SetFilter("chat_id>?", maxChatId)
		w.SetFilter("status IN(?)", []pb_enum.CHAT_STATUS{pb_enum.CHAT_STATUS_NORMAL, pb_enum.CHAT_STATUS_MUTE, pb_enum.CHAT_STATUS_BANNED})
		w.SetLimit(int32(limit))
		w.SetSort("chat_id ASC")
		list, err = s.chatMemberRepo.ChatMemberStatusList(w)
		if err != nil {
			return
		}
		if len(list) > 0 {
			if allStatus == nil {
				allStatus = list
			} else {
				allStatus = append(allStatus, list...)
			}
		}
		if len(list) < limit {
			break
		}
		maxChatId = list[len(list)-1].ChatId
	}

	// 2 更新users
	u.SetFilter("uid = ?", uid)
	u.Set("server_id", serverId)
	err = s.userRepo.UpdateUser(u)
	if err != nil {
		return
	}
	if len(allStatus) == 0 {
		return
	}

	xants.Submit(func() {
		// 3 更新hash
		var (
			keys      = make([]string, len(allStatus))
			vals      = make([]string, len(allStatus))
			index     int
			chat      *do.ChatMemberStatus
			valPrefix = fmt.Sprintf("%d,", serverId)
			uidStr    = utils.ToString(uid)
			err       error
		)
		for index, chat = range allStatus {
			keys[index] = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(chat.ChatId)
			vals[index] = valPrefix + utils.IntToStr(int(chat.Status))
		}
		err = s.chatMemberCache.HSetDistChatMembers(keys, uidStr, vals)
		if err != nil {
			xlog.Warn(err.Error())
			// 消息队列
			msg := &do.KeysFieldValues{keys, uidStr, vals}
			_, _, err = s.producer.Push(msg, constant.CONST_MSG_KEY_CACHE_ON_OFF_LINE)
			if err != nil {
				xlog.Warn(err.Error())
			}
		}
	})
	return
}

// 以下代码被弃用
/*
func (s *chatMemberService) cacheChatMember(list []*do.ChatMemberInfo) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		threads     = runtime.NumCPU() * 2
		ctx, cancel = context.WithCancel(context.Background())
		memberCh    = make(chan *do.ChatMemberInfo, 0)
		completeCh  = make(chan int, threads)
		i           int
		code        int
	)
	s.consumerMessage(ctx, cancel, memberCh, completeCh, threads)
	xants.Submit(func() {
		s.productionMessage(cancel, list, memberCh)
	})
	for i = 0; i < threads; i++ {
		code = <-completeCh
		if code != 0 {
			err = ERR_CHAT_MEMBER_CHCHE_MEMBER_FAILED
			break
		}
	}
	return
}

func (s *chatMemberService) productionMessage(cancel context.CancelFunc, list []*do.ChatMemberInfo, memberCh chan *do.ChatMemberInfo) {
	var (
		member *do.ChatMemberInfo
	)
	for _, member = range list {
		memberCh <- member
	}
	cancel()
}

func (s *chatMemberService) consumerMessage(ctx context.Context, cancel context.CancelFunc, memberCh chan *do.ChatMemberInfo, completeCh chan int, threads int) {
	var (
		i   int
		key string
	)
	for i = 0; i < threads; i++ {
		xants.Submit(func() {
			var (
				err  error
				code int
			)
			defer func() {
				code = 0
				if err != nil {
					code = 1
				}
				completeCh <- code
			}()
			var (
				member *do.ChatMemberInfo
				ok     bool
				val    string
			)
			for {
				select {
				case <-ctx.Done():
					return
				case member, ok = <-memberCh:
					if ok == false {
						return
					}
					key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(member.ChatId)
					// 0:ServerId, 1:Platform, 2:Uid, 3:Status
					val = fmt.Sprintf("%d,%d", member.ServerId, member.Status)
					err = xredis.HMSet(key, map[string]interface{}{utils.Int64ToStr(member.Uid): val})
					if err != nil {
						cancel()
						return
					}
				}
			}
		})
	}
	return
}

*/
