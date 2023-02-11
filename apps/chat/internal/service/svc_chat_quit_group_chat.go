package service

import (
	"context"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *chatService) QuitGroupChat(ctx context.Context, req *pb_chat.QuitGroupChatReq) (resp *pb_chat.QuitGroupChatResp, _ error) {
	resp = new(pb_chat.QuitGroupChatResp)
	var (
		u   = entity.NewMysqlUpdate()
		err error
	)
	u.SetFilter("chat_id=?", req.ChatId)
	u.SetFilter("uid=?", req.Uid)
	u.SetFilter("deleted_ts=?", 0)
	u.Set("status", int32(pb_enum.CHAT_STATUS_QUITTED))
	u.Set("deleted_ts", utils.NowMilli())
	_, err = s.removeChatMember(u, req.ChatId, []int64{req.Uid}, pb_enum.CHAT_TYPE_GROUP)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED, err.Error())
		return
	}

	// 推送成员退出
	xants.Submit(func() {
		s.quitGroupChatMessage(
			req.ChatId,
			[]int64{req.Uid},
			pb_enum.SUB_TOPIC_CHAT_QUIT_GROUP_CHAT,
			pb_enum.MSG_TYPE_QUIT_GROUP_CHAT)
	})
	return
}

func (s *chatService) quitGroupChatMessage(chatId int64, uidList []int64, subTopic pb_enum.SUB_TOPIC, msgType pb_enum.MSG_TYPE) {
	var (
		seqId      int64
		nowMilli   = utils.NowMilli()
		w          = entity.NewNormalWhere()
		memberList []*pb_chat_member.ChatMemberBasicInfo
		msg        *pb_msg.SrvChatMessage
		inbox      *pb_mq.InboxMessage
		jsonStr    string
		err        error
	)
	if seqId, err = s.chatMessageCache.IncrSeqID(chatId); err != nil {
		return
	}
	msg = &pb_msg.SrvChatMessage{
		SrvMsgId:        xsnowflake.NewSnowflakeID(),
		CliMsgId:        xsnowflake.NewSnowflakeID(),
		SenderId:        0,
		SenderPlatform:  0,
		SenderName:      "",
		SenderAvatarKey: "",
		ChatId:          chatId,
		ChatType:        pb_enum.CHAT_TYPE_GROUP,
		SeqId:           seqId,
		MsgFrom:         pb_enum.MSG_FROM_SYSTEM,
		MsgType:         msgType,
		Body:            nil,
		Status:          0,
		SentTs:          nowMilli,
		SrvTs:           nowMilli,
	}
	w.SetFilter("chat_id=?", chatId)
	w.SetFilter("uid IN(?)", uidList)
	memberList, err = s.chatMemberRepo.ChatMemberBasicInfoList(w)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED, err.Error())
		return
	}
	if len(memberList) == 0 {
		return
	}
	jsonStr, err = utils.Marshal(memberList)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_PROTOCOL_MARSHAL_ERR, ERROR_CHAT_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	msg.Body = []byte(jsonStr)

	inbox = &pb_mq.InboxMessage{
		Topic:    pb_enum.TOPIC_CHAT,
		SubTopic: subTopic,
		Msg:      msg,
	}
	// 将消息推送到kafka消息队列
	_, _, err = s.producer.EnQueue(inbox, constant.CONST_MSG_KEY_MSG)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_ENQUEUE_FAILED, ERROR_CHAT_ENQUEUE_FAILED, err.Error())
		return
	}
}
