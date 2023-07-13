package service

import (
	"lark/domain/po"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_msg"
)

func (s *chatMessageService) GetHistoryMessages(req *pb_chat_msg.GetChatMessagesReq) (list []*po.Message, err error) {
	// 从mysql中获取消息
	var (
		w = entity.NewNormalQuery()
	)
	w.Limit = int(req.Limit)
	w.SetFilter("chat_id = ?", req.ChatId)
	if req.New {
		w.Sort = "seq_id ASC"
		w.SetFilter("seq_id > ?", req.SeqId)
	} else {
		w.Sort = "seq_id DESC"
		w.SetFilter("seq_id < ?", req.SeqId)
	}
	list, err = s.chatMessageRepo.HistoryMessages(w)
	if err != nil {
		return
	}
	return
}

func (s *chatMessageService) GetHistoryMessageList(req *pb_chat_msg.GetChatMessageListReq) (list []*po.Message, err error) {
	var (
		w = entity.NewNormalQuery()
	)
	//w.Sort = "seq_id ASC"
	w.SetFilter("chat_id = ?", req.ChatId)
	w.SetFilter("seq_id in(?)", req.SeqIds)
	list, err = s.chatMessageRepo.HistoryMessages(w)
	if err != nil {
		return
	}
	return
}
