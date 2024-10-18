package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/proto/pb_obj"
	"sync"
)

func (s *distService) messageOperation(buf []byte) (err error) {
	var (
		req           = new(pb_msg.MessageOperationReq)
		serverMembers map[int64][]*pb_obj.Int64Array
		wg            = new(sync.WaitGroup)
	)
	if err = proto.Unmarshal(buf, req); err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_UNMARSHAL_ERR, ERROR_DIST_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	slots := int(constant.MAX_CHAT_SLOT)
	if req.Operation.ChatType == pb_enum.CHAT_TYPE_PRIVATE {
		slots = 0
	}
	for i := 0; i < slots; i++ {
		wg.Add(1)
		go func(slot int) {
			defer wg.Done()
			serverMembers = s.getChatMembers(req.Operation.ChatId, slot)
			s.sendMessageOperation(req, serverMembers)
		}(i)
	}
	wg.Wait()
	return
}

func (s *distService) sendMessageOperation(opnReq *pb_msg.MessageOperationReq, distMembers map[int64][]*pb_obj.Int64Array) {
	if len(distMembers) == 0 {
		return
	}
	var (
		serverId int64
		body     []byte
		wg       = new(sync.WaitGroup)
		err      error
	)
	body, err = proto.Marshal(opnReq.Operation)
	if err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_MARSHAL_ERR, ERROR_DIST_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	for serverId, _ = range distMembers {
		msgReq := &pb_gw.SendTopicMessageReq{
			Topic:          opnReq.Topic,
			SubTopic:       opnReq.SubTopic,
			Members:        distMembers[serverId],
			SenderId:       opnReq.Operation.SenderId,
			SenderPlatform: opnReq.Operation.Platform,
			Body:           body,
		}
		s.asyncSendMessage(wg, msgReq, serverId, constant.CONST_MSG_KEY_PUSH_OPERATION)
	}
	wg.Wait()
	return
}
