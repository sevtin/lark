package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *messageHistoryService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_MSG:
		err = s.SaveMessage(msg)
		return
	case constant.CONST_MSG_KEY_OPERATION:
		err = s.MessageOperation(msg)
		return
	default:
		return
	}
}

func (s *messageHistoryService) SaveMessage(msg []byte) (err error) {
	var (
		req     = new(pb_mq.InboxMessage)
		message = new(po.Message)
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		// 丢弃无法解析的数据
		err = nil
		return
	}
	if req.Msg.SrvMsgId == 0 {
		return
	}
	copier.Copy(message, req.Msg)
	message.Body = utils.MsgBodyToStr(req.Msg.MsgType, req.Msg.Body)

	// 1、消息缓存
	xants.Submit(func() {
		s.chatMessageCache.SetConvoMessage(message)
	})

	// 2、消息入库
	if err = s.chatMessageRepo.Create(message); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_INSERT_MESSAGE_FAILED, ERROR_MSG_HISTORY_INSERT_MESSAGE_FAILED, err.Error())
		switch err.(type) {
		case *mysql.MySQLError:
			if err.(*mysql.MySQLError).Number == constant.ERROR_CODE_MYSQL_DUPLICATE_ENTRY {
				err = nil
				return
			}
		}
		return
	}
	return
}

func (s *messageHistoryService) MessageOperation(msg []byte) (err error) {
	var (
		req      = new(pb_msg.MessageOperationReq)
		u        = entity.NewMysqlUpdate()
		message  *po.Message
		nowMilli = utils.NowMilli()
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	message, err = s.GetMessage(req.Operation.ChatId, req.Operation.SeqId)
	if err != nil {
		return
	}
	if message.SrvMsgId == 0 {
		return
	}

	// 1、更新缓存
	message.Status = int(req.Operation.Opn)
	message.UpdatedTs = nowMilli
	err = s.chatMessageCache.SetChatMessage(message)
	if err != nil {
		return
	}

	// 2、更新 message
	u.SetFilter("srv_msg_id=?", req.Operation.SrvMsgId)
	u.Set("status", req.Operation.Opn)
	err = s.chatMessageRepo.UpdateMessage(u)
	if err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_UPDATE_VALUE_FAILED, ERROR_MSG_HISTORY_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	return
}

func (s *messageHistoryService) GetMessage(chatId int64, seqId int64) (message *po.Message, err error) {
	message = new(po.Message)
	var (
		w = entity.NewNormalWhere()
	)
	message, err = s.chatMessageCache.GetChatMessage(chatId, seqId)
	if message.SrvMsgId > 0 {
		return
	}

	if err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_REDIS_GET_FAILED, ERROR_MSG_HISTORY_REDIS_GET_FAILED, err.Error())
	}
	w.SetFilter("chat_id=?", chatId)
	w.SetFilter("seq_id=?", seqId)
	message, err = s.chatMessageRepo.Message(w)
	if err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_QUERY_DB_FAILED, ERROR_MSG_HISTORY_QUERY_DB_FAILED, err.Error())
		return
	}
	return
}
