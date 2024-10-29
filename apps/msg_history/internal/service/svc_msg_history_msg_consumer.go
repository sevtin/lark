package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *messageHistoryService) MessageHandler(msg []byte, msgKey string) (err error) {
	inbox := new(pb_mq.InboxMessage)
	if err = proto.Unmarshal(msg, inbox); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	switch inbox.Topic {
	case pb_enum.TOPIC_CHAT, pb_enum.TOPIC_RED_ENVELOPE:
		err = s.SaveMessage(inbox.Body)
	case pb_enum.TOPIC_MSG_OPR:
		err = s.MessageOperation(inbox.Body)
	default:
		return
	}
	return
}

func (s *messageHistoryService) SaveMessage(msg []byte) (err error) {
	var (
		chatMsg = new(pb_msg.SrvChatMessage)
		message = new(po.Message)
	)
	if err = proto.Unmarshal(msg, chatMsg); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	if chatMsg.SrvMsgId == 0 {
		return
	}
	copier.Copy(message, chatMsg)
	message.Body = utils.MsgBodyToStr(chatMsg.MsgType, chatMsg.Body)

	// 1、消息入库
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

	// 2、消息缓存
	s.chatMessageCache.SetConvoMessage(message)
	return
}

func (s *messageHistoryService) MessageOperation(msg []byte) (err error) {
	var (
		operation = new(pb_msg.MessageOperation)
		u         = entity.NewMysqlUpdate()
		message   *po.Message
		nowTs     = utils.NowUnix()
	)
	if err = proto.Unmarshal(msg, operation); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	message, err = s.GetMessage(operation.ChatId, operation.SeqId)
	if err != nil {
		return
	}
	if message.SrvMsgId == 0 {
		return
	}
	message.Status = int(operation.Opn)
	message.UpdatedTs = nowTs
	// 1、更新 message
	u.SetFilter("chat_id=?", operation.ChatId)
	u.SetFilter("seq_id=?", operation.SeqId)
	u.Set("status", operation.Opn)
	err = s.chatMessageRepo.UpdateMessage(u)
	if err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_UPDATE_VALUE_FAILED, ERROR_MSG_HISTORY_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	// 2、更新缓存
	err = s.chatMessageCache.SetChatMessage(message)
	if err != nil {
		return
	}
	return
}

func (s *messageHistoryService) GetMessage(chatId int64, seqId int64) (message *po.Message, err error) {
	message = new(po.Message)
	var (
		w = entity.NewNormalQuery()
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
