package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"sync"
)

func (s *distService) messageOperation(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, buf []byte) (err error) {
	var (
		operation     = new(pb_msg.MessageOperation)
		serverMembers map[int64][]*pb_obj.Int64Array
		wg            = new(sync.WaitGroup)
	)
	if err = proto.Unmarshal(buf, operation); err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_UNMARSHAL_FAILED, ERROR_DIST_PROTOCOL_UNMARSHAL_FAILED, err.Error())
		return
	}
	slots := int(constant.MAX_CHAT_SLOT)
	if operation.ChatType == pb_enum.CHAT_TYPE_PRIVATE {
		slots = 0
	}
	for i := 0; i < slots; i++ {
		wg.Add(1)
		go func(slot int) {
			defer wg.Done()
			serverMembers = s.getChatMembers(operation.ChatId, slot)
			s.sendMessageOperation(topic, subTopic, operation, buf, serverMembers)
		}(i)
	}
	wg.Wait()
	return
}

func (s *distService) sendMessageOperation(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, operation *pb_msg.MessageOperation, body []byte, distMembers map[int64][]*pb_obj.Int64Array) {
	if len(distMembers) == 0 {
		return
	}
	var (
		serverId int64
		wg       = new(sync.WaitGroup)
	)
	for serverId, _ = range distMembers {
		msgReq := &pb_gw.SendTopicMessageReq{
			Topic:          topic,
			SubTopic:       subTopic,
			Members:        distMembers[serverId],
			SenderId:       operation.SenderId,
			SenderPlatform: operation.Platform,
			Body:           body,
		}
		s.asyncSendMessage(wg, msgReq, serverId, constant.CONST_MSG_KEY_PUSH_OPERATION+utils.GetChatPartition(operation.ChatId))
	}
	wg.Wait()
	return
}
