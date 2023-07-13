package service

import (
	"context"
	"fmt"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
)

func (s *chatMemberService) GetDistMemberList(ctx context.Context, req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp, _ error) {
	resp = &pb_chat_member.GetDistMemberListResp{}
	var (
		w       = entity.NewMysqlQuery()
		count   int
		lastUid int64
		members []*pb_chat_member.DistMember
		member  *pb_chat_member.DistMember
		maps    = make(map[string]string)
		err     error
	)

	for {
		var (
			value string
		)
		w.Normal()
		w.SetFilter("m.chat_id = ?", req.ChatId)
		w.SetFilter("m.deleted_ts = ?", 0)
		w.SetFilter("m.status IN(?)", []pb_enum.CHAT_STATUS{pb_enum.CHAT_STATUS_NORMAL, pb_enum.CHAT_STATUS_MUTE, pb_enum.CHAT_STATUS_BANNED})
		w.SetFilter("m.uid > ?", lastUid)
		w.SetSort("m.uid ASC")
		w.SetLimit(1000)

		members, err = s.chatMemberRepo.DistMemberList(w)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
			return
		}
		count = len(members)
		if count == 0 {
			break
		}
		for _, member = range members {
			// 0:ServerId, 1:Platform, 2:Uid, 3:Status
			value = fmt.Sprintf("%d,%d", member.ServerId, member.Status)
			maps[utils.Int64ToStr(member.Uid)] = value
		}
		if count < w.Limit {
			break
		}
		lastUid = members[count-1].Uid
	}
	resp.Members = maps
	if len(maps) == 0 {
		return
	}
	err = s.chatMemberCache.HMSetChatMembers(req.ChatId, maps)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED, ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED, err.Error())
		return
	}
	return
}

/*
func (s *chatMemberService) cacheDistMember(list []*pb_chat_member.DistMember, chatId int64) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		key     string
		member  *pb_chat_member.DistMember
		jsonStr string
		members = map[string]interface{}{}
	)
	for _, member = range list {
		jsonStr, _ = utils.Marshal(member)
		members[utils.Int64ToStr(member.Uid)] = jsonStr
	}
	key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(chatId)
	err = xredis.HMSet(key, members)
	return
}
*/
