package service

import (
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/utils"
	"strconv"
)

func (s *chatMessageService) GetCacheMessages(req *pb_chat_msg.GetChatMessagesReq, maxSeqId int64) (list []*po.Message, next bool, err error) {
	var (
		msgCount int
	)
	list = s.GetCacheChatMessages(req, int(maxSeqId))
	msgCount = len(list)
	if msgCount == int(req.Limit) {
		return
	}
	if msgCount > 0 {
		req.SeqId = list[msgCount-1].SeqId
		if req.SeqId >= maxSeqId {
			return
		}
		req.Limit -= int32(msgCount)
	}
	next = true
	return
}

func (s *chatMessageService) GetCacheChatMessages(req *pb_chat_msg.GetChatMessagesReq, max int) (list []*po.Message) {
	list = make([]*po.Message, 0)
	var (
		minSeqId int
		maxSeqId int
		seqId    int
		dv       int
		index    int
		key      string
		values   []interface{}
		val      interface{}
		err      error
	)
	if req.New == true {
		minSeqId = int(req.SeqId) + 1
		maxSeqId = minSeqId + int(req.Limit) - 1
	} else {
		maxSeqId = int(req.SeqId) - 1
		minSeqId = maxSeqId - int(req.Limit) + 1
	}
	if minSeqId < 0 {
		minSeqId = 0
	}
	if maxSeqId < 0 {
		maxSeqId = 0
	}

	if maxSeqId > max {
		maxSeqId = max
	}
	dv = maxSeqId - minSeqId
	if dv <= 0 {
		return
	}
	keys := make([]string, dv+1)
	for index = 0; index <= dv; index++ {
		if req.New == true {
			seqId = minSeqId + index
		} else {
			seqId = maxSeqId - index
		}
		key = s.cfg.Redis.Prefix + constant.RK_SYNC_MSG_CACHE + utils.Int64ToStr(req.ChatId) + ":" + strconv.Itoa(seqId)
		keys[index] = key
	}
	values, err = s.chatMessageCache.MGetMessages(keys...)
	if err != nil {
		return
	}
	for _, val = range values {
		msg := &po.Message{}
		if val == nil {
			xlog.Warn(ERROR_CODE_CHAT_MSG_REDIS_GET_FAILED, ERROR_CHAT_MSG_REDIS_GET_FAILED, err.Error())
			break
		}
		utils.Unmarshal(val.(string), msg)
		list = append(list, msg)
	}
	return
}

func (s *chatMessageService) SaveCacheChatMessageCache(list []*po.Message) {
	if len(list) == 0 {
		return
	}
	var (
		index int
	)
	for index, _ = range list {
		s.chatMessageCache.SetChatMessage(list[index])
	}
}

func (s *chatMessageService) cacheChatMessage(msg *po.Message) {
	xants.Submit(func() {
		s.chatMessageCache.SetChatMessage(msg)
	})
}

/*
---------------------- 分割线 ----------------------
*/

func (s *chatMessageService) GetCacheChatMessageList(req *pb_chat_msg.GetChatMessageListReq) (list []*po.Message, err error) {
	var (
		index   int
		seqId   int64
		key     string
		keys    = make([]string, len(req.SeqIds))
		seqMaps = map[int64]int{}
		values  []interface{}
		val     interface{}
	)
	for index, seqId = range req.SeqIds {
		key = s.cfg.Redis.Prefix + constant.RK_SYNC_MSG_CACHE + utils.Int64ToStr(req.ChatId) + ":" + utils.Int64ToStr(seqId)
		keys[index] = key
		seqMaps[seqId] = 0
	}
	values, err = s.chatMessageCache.MGetMessages(keys...)
	if err != nil {
		return
	}
	for _, val = range values {
		if val == nil {
			continue
		}
		msg := &po.Message{}
		utils.Unmarshal(val.(string), msg)
		list = append(list, msg)
		delete(seqMaps, msg.SeqId)
	}
	req.SeqIds = make([]int64, len(seqMaps))
	index = 0
	for seqId, _ = range seqMaps {
		req.SeqIds[index] = seqId
		index++
	}
	return
}
