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
	switch msgKey {
	case constant.CONST_MSG_KEY_MSG:
		err = s.SaveMessage(msg)
		return
	case constant.CONST_MSG_KEY_RECALL:
		err = s.MessageRecall(msg)
		return
	default:
		return
	}
}

func (s *messageHotService) SaveMessage(msg []byte) (err error) {
	var (
		req     = new(pb_mq.InboxMessage)
		message = new(po.Message)
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, err.Error())
		// 丢弃无法解析的数据
		err = nil
	}
	// 消息入库
	copier.Copy(message, req.Msg)
	message.Body = utils.MsgBodyToStr(req.Msg.MsgType, req.Msg.Body)
	message.UpdatedTs = utils.NowMilli()
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
		req = new(pb_msg.MessageOperationReq)
		u   = entity.NewMongoUpdate()
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	u.SetFilter("srv_msg_id", req.Operation.SrvMsgId)
	u.SetFilter("sender_id", req.Operation.SenderId)
	u.Set("status", pb_enum.MSG_OPERATION_RECALL)
	if err = s.messageHotRepo.Update(u); err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}
