package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"lark/domain/po"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_msg"
)

func (s *chatMessageService) GetHotMessages(req *pb_chat_msg.GetChatMessagesReq, maxSeqId int64) (list []*po.Message, next bool, err error) {
	// 从mongo中获取消息
	var (
		w        = entity.NewMongoWhere()
		msgCount int
	)
	w.SetLimit(int64(req.Limit))
	w.SetSort("seq_id", true)
	w.SetFilter("chat_type", req.ChatType)
	w.SetFilter("chat_id", req.ChatId)

	if req.SeqId > 0 {
		if req.New {
			//大于($gt)
			w.SetFilter("seq_id", bson.M{"$gt": req.SeqId})
		} else {
			//小于($lt)
			w.SetFilter("seq_id", bson.M{"$lt": req.SeqId})
			//大于或等于($gte)
			w.SetFilter("seq_id", bson.M{"$gte": req.SeqId - int64(req.Limit)})
		}
	}
	list, err = s.messageHotRepo.Messages(w)
	msgCount = len(list)
	if msgCount == int(req.Limit) {
		return
	}
	if msgCount > 0 {
		if req.New == true {
			req.SeqId = list[msgCount-1].SeqId
		} else {
			req.SeqId = list[0].SeqId
		}
		if req.SeqId >= maxSeqId {
			return
		}
		req.Limit -= int32(msgCount)
	}
	next = true
	return
}
