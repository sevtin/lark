package service

import (
	"context"
	"fmt"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

func (s *chatMemberService) GetDistMemberList(ctx context.Context, req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp, _ error) {
	resp = &pb_chat_member.GetDistMemberListResp{List: make([]string, 0)}
	var (
		w       = entity.NewMysqlWhere()
		count   int
		lastUid int64
		members []*pb_chat_member.DistMember
		member  *pb_chat_member.DistMember
		maps    = make(map[string]interface{})
		err     error
	)

	for {
		var (
			values []string
			index  int
			value  string
		)
		w.Reset()
		w.SetFilter("chat_id = ?", req.ChatId)
		w.SetFilter("uid > ?", lastUid)
		w.SetSort("uid ASC")
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
		values = make([]string, count)
		for index, member = range members {
			// 0:ServerId, 1:Platform, 2:Uid, 3:Status
			value = fmt.Sprintf("%d,%d,%d", member.ServerId, member.Uid, member.Status)
			values[index] = value
			maps[utils.Int64ToStr(member.Uid)] = value
		}
		resp.List = append(resp.List, values...)
		if count < w.Limit {
			break
		}
		lastUid = members[count-1].Uid
	}
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
