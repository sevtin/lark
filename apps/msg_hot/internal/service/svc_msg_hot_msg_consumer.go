package service

import (
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *messageHotService) MessageHandler(msg []byte, msgKey string) (err error) {
	inbox := new(pb_mq.InboxMessage)
	if err = proto.Unmarshal(msg, inbox); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, err.Error())
		return
	}
	switch inbox.Topic {
	case pb_enum.TOPIC_CHAT, pb_enum.TOPIC_RED_ENVELOPE:
		err = s.SaveMessage(inbox.Body)
	case pb_enum.TOPIC_MSG_OPR:
		err = s.MessageRecall(inbox.Body)
	default:
		return
	}
	return
}

func (s *messageHotService) SaveMessage(msg []byte) (err error) {
	var (
		chatMsg = new(pb_msg.SrvChatMessage)
		message = new(po.Message)
	)
	if err = proto.Unmarshal(msg, chatMsg); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, err.Error())
		return
	}
	// 消息入库
	copier.Copy(message, chatMsg)
	message.Body = utils.MsgBodyToStr(chatMsg.MsgType, chatMsg.Body)
	message.UpdatedTs = utils.NowUnix()
	if err = s.messageHotRepo.Create(message); err != nil {
		xlog.Warn(err.Error())
		switch err.(type) {
		case mongo.WriteException:
			if len(err.(mongo.WriteException).WriteErrors) > 0 {
				if err.(mongo.WriteException).WriteErrors[0].Code == constant.ERROR_CODE_MONGOL_DUPLICATE_ENTRY {
					err = nil
				}
			}
		}
		return
	}
	return
}

func (s *messageHotService) MessageRecall(msg []byte) (err error) {
	var (
		operation = new(pb_msg.MessageOperation)
		u         = entity.NewMongoUpdate()
	)
	if err = proto.Unmarshal(msg, operation); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_FAILED, err.Error())
		return
	}
	u.SetFilter("srv_msg_id", operation.SrvMsgId)
	u.SetFilter("sender_id", operation.SenderId)
	u.Set("status", pb_enum.MSG_OPERATION_RECALL)
	if err = s.messageHotRepo.Update(u); err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}
