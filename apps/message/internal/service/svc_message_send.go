package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *messageService) SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error) {
	resp = new(pb_msg.SendChatMessageResp)
	var (
		inbox = &pb_mq.InboxMessage{
			Topic:    req.Topic,
			SubTopic: req.SubTopic,
			Msg:      new(pb_msg.SrvChatMessage),
		}
		senderInfo *pb_chat_member.ChatMemberInfo
		seqId      int64
		result     string
		ok         bool
		err        error
	)
	// 1、参数校验
	if err = s.validate.Struct(req.Msg); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR)
		xlog.Warn(ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR, err.Error())
		return
	}
	if inbox.Msg.AssocId, err = s.verifyMessage(req); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR)
		xlog.Warn(ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR, err.Error())
		return
	}
	// 2、重复消息校验
	result, ok = s.chatMessageCache.RepeatMessageVerify(s.cfg.Redis.Prefix, req.Msg.ChatId, req.Msg.CliMsgId)
	if ok == false {
		resp.Set(ERROR_CODE_MESSAGE_VALIDATOR_ERR, result)
		xlog.Warn(ERROR_CODE_MESSAGE_VALIDATOR_ERR, result)
		return
	}

	// 3、获取发送者信息
	senderInfo, err = s.getSenderInfo(req.Msg.ChatId, req.Msg.SenderId, resp)
	if err != nil {
		resp.Set(ERROR_CODE_MESSAGE_GET_SENDER_INFO_FAILED, err.Error())
		xlog.Warn(ERROR_CODE_MESSAGE_GET_SENDER_INFO_FAILED, ERROR_MESSAGE_GET_SENDER_INFO_FAILED, err.Error())
		return
	}
	if senderInfo.Uid == 0 {
		if resp.Code > 0 {
			xlog.Warn(resp.Code, resp.Msg)
		} else {
			xlog.Warn(ERROR_CODE_MESSAGE_GET_SENDER_INFO_FAILED, ERROR_MESSAGE_GET_SENDER_INFO_FAILED)
		}
		return
	}

	// 4、补充消息内容
	if seqId, err = s.chatMessageCache.IncrSeqID(req.Msg.ChatId); err != nil {
		resp.Set(ERROR_CODE_MESSAGE_INCR_SEQ_ID_FAILED, ERROR_MESSAGE_INCR_SEQ_ID_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_INCR_SEQ_ID_FAILED, ERROR_MESSAGE_INCR_SEQ_ID_FAILED, err.Error())
		return
	}
	copier.Copy(inbox.Msg, req.Msg)
	inbox.Msg.SrvMsgId = xsnowflake.NewSnowflakeID()
	inbox.Msg.ChatType = senderInfo.ChatType
	inbox.Msg.SeqId = seqId
	inbox.Msg.SrvTs = utils.NowUnix()
	inbox.Msg.MsgFrom = pb_enum.MSG_FROM_USER
	inbox.Msg.SenderName = senderInfo.Alias
	inbox.Msg.SenderAvatar = senderInfo.MemberAvatar

	// 5、将消息推送到kafka消息队列
	if s.producer == nil {
		resp.Set(ERROR_CODE_MESSAGE_PRODUCER_IS_NULL, ERROR_MESSAGE_PRODUCER_IS_NULL)
		xlog.Warn(ERROR_CODE_MESSAGE_PRODUCER_IS_NULL, ERROR_MESSAGE_PRODUCER_IS_NULL)
		return
	}
	_, _, err = s.producer.EnQueue(inbox, constant.CONST_MSG_KEY_MSG)
	if err != nil {
		resp.Set(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED, err.Error())
		return
	}
	chatSeq := &pb_msg.ChatSeq{
		ChatId:   inbox.Msg.ChatId,
		SeqId:    seqId,
		SrvTs:    inbox.Msg.SrvTs,
		SenderId: inbox.Msg.SenderId,
		MsgFrom:  inbox.Msg.MsgFrom,
	}
	_, _, err = s.seqProducer.EnQueue(chatSeq, constant.CONST_MSG_KEY_CHAT_SEQ)
	if err != nil {
		xlog.Warn(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED, err.Error())
	}
	return
}

func (s *messageService) getSenderInfo(chatId int64, uid int64, resp *pb_msg.SendChatMessageResp) (info *pb_chat_member.ChatMemberInfo, err error) {
	var (
		req   *pb_chat_member.GetChatMemberInfoReq
		reply *pb_chat_member.GetChatMemberInfoResp
	)
	info, err = s.chatMemberCache.GetChatMemberInfo(chatId, uid)
	if err != nil {
		xlog.Warn(ERROR_CODE_MESSAGE_REDIS_GET_FAILED, ERROR_MESSAGE_REDIS_GET_FAILED, err.Error())
	}
	if info.Uid > 0 {
		err = s.authentication(info)
		if err != nil {
			return
		}
		return
	}
	req = &pb_chat_member.GetChatMemberInfoReq{
		ChatId: chatId,
		Uid:    uid,
	}
	reply = s.chatMemberClient.GetChatMemberInfo(req)
	if reply == nil {
		xlog.Warn(ERROR_CODE_MESSAGE_GRPC_SERVICE_FAILURE, ERROR_MESSAGE_GRPC_SERVICE_FAILURE)
		resp.Set(ERROR_CODE_MESSAGE_GRPC_SERVICE_FAILURE, ERROR_MESSAGE_GRPC_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		xlog.Warn(reply.Code, reply.Msg)
		resp.Set(reply.Code, reply.Msg)
		return
	}
	err = s.authentication(reply.Info)
	if err != nil {
		return
	}
	info = reply.Info
	return
}

func (s *messageService) authentication(info *pb_chat_member.ChatMemberInfo) (err error) {
	switch pb_enum.CHAT_STATUS(info.Status) {
	case pb_enum.CHAT_STATUS_QUITTED, pb_enum.CHAT_STATUS_DELETED, pb_enum.CHAT_STATUS_REMOVED:
		// 非联 Chat 成员
		err = ERROR_MESSAGE_NON_CHAT_MEMBER_ERR
	case pb_enum.CHAT_STATUS_NON_CONTACT:
		// 非联系人
		err = ERROR_MESSAGE_NON_CONTACT_ERR
	}
	return
}
