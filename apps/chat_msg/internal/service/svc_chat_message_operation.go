package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *chatMessageService) MessageOperation(ctx context.Context, req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp, err error) {
	resp = new(pb_chat_msg.MessageOperationResp)
	var (
		message  *po.Message
		nowMilli = utils.NowMilli()
		opnReq   = &pb_msg.MessageOperationReq{
			Topic:     pb_enum.TOPIC_CHAT,
			SubTopic:  pb_enum.SUB_TOPIC_CHAT_OPERATION,
			Operation: &pb_msg.MessageOperation{},
		}
	)
	message, err = s.GetMessage(req.ChatId, req.SeqId)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_QUERY_DB_FAILED, ERROR_CHAT_MSG_QUERY_DB_FAILED)
		return
	}
	// 1、超过10分钟无法测回
	if nowMilli-message.SrvTs > constant.CONST_MILLISECOND_10_MINUTES {
		resp.Set(ERROR_CODE_CHAT_MSG_BEYOND_OPERABLE_TIME, ERROR_CHAT_MSG_BEYOND_OPERABLE_TIME)
		xlog.Warn(ERROR_CODE_CHAT_MSG_BEYOND_OPERABLE_TIME, ERROR_CHAT_MSG_BEYOND_OPERABLE_TIME)
		return
	}
	switch pb_enum.MSG_OPERATION(message.Status) {
	case pb_enum.MSG_OPERATION_RECALL:
		// 重复操作
		return
	}
	// 2、成员身份判断
	if message.SenderId != req.SenderId {
		// 无权操作
		return
	}

	// 3、消息推送
	copier.Copy(opnReq.Operation, req)
	opnReq.Operation.SrvMsgId = message.SrvMsgId
	_, _, err = s.producer.EnQueue(opnReq, constant.CONST_MSG_KEY_OPERATION)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_ENQUEUE_FAILED, ERROR_CHAT_MSG_ENQUEUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_MSG_ENQUEUE_FAILED, ERROR_CHAT_MSG_ENQUEUE_FAILED, err.Error())
		return
	}
	return
}

func (s *chatMessageService) GetMessage(chatId int64, seqId int64) (message *po.Message, err error) {
	var (
		w = entity.NewNormalWhere()
	)
	message, _ = s.chatMessageCache.GetChatMessage(chatId, seqId)
	if message.SrvMsgId > 0 {
		return
	}
	w.SetFilter("chat_id=?", chatId)
	w.SetFilter("seq_id=?", seqId)
	message, err = s.chatMessageRepo.Message(w)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MSG_QUERY_DB_FAILED, ERROR_CHAT_MSG_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
