package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"sync"
)

func (s *distService) sendChatMessage(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, buf []byte) (err error) {
	var (
		chatMsg = new(pb_msg.SrvChatMessage)
	)
	if err = proto.Unmarshal(buf, chatMsg); err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_UNMARSHAL_FAILED, ERROR_DIST_PROTOCOL_UNMARSHAL_FAILED, err.Error())
		return
	}
	s.messageDistribution(topic, subTopic, chatMsg, buf)
	return
}

func (s *distService) messageDistribution(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, msg *pb_msg.SrvChatMessage, body []byte) {
	wg := new(sync.WaitGroup)
	slots := int(constant.MAX_CHAT_SLOT)
	if msg.ChatType == pb_enum.CHAT_TYPE_PRIVATE {
		slots = 1
	}
	for i := 0; i < slots; i++ {
		wg.Add(1)
		go func(slot int) {
			defer wg.Done()
			distMembers := s.getChatMembers(msg.ChatId, slot)
			s.sendMessage(topic, subTopic, msg, body, distMembers)
		}(i)
	}
	wg.Wait()
}

func (s *distService) sendMessage(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, msg *pb_msg.SrvChatMessage, body []byte, distMembers map[int64][]*pb_obj.Int64Array) {
	if msg.SrvMsgId == 0 {
		return
	}
	if len(distMembers) == 0 {
		return
	}
	var (
		wg = new(sync.WaitGroup)
	)
	for serverId, _ := range distMembers {
		msgReq := &pb_gw.SendTopicMessageReq{
			Topic:          topic,
			SubTopic:       subTopic,
			Members:        distMembers[serverId],
			SenderId:       msg.SenderId,
			SenderPlatform: msg.SenderPlatform,
			Body:           body,
		}
		s.asyncSendMessage(wg, msgReq, serverId, constant.CONST_MSG_KEY_PUSH_ONLINE+utils.GetChatPartition(msg.ChatId))
	}
	wg.Wait()
	return
}

func (s *distService) asyncSendMessage(wg *sync.WaitGroup, req *pb_gw.SendTopicMessageReq, serverId int64, key string) {
	wg.Add(1)
	xants.Submit(func() {
		defer wg.Done()
		var (
			resp *pb_gw.SendTopicMessageResp
			err  error
		)
		client := s.getClient(serverId)
		if client == nil {
			xlog.Warn(ERROR_CODE_DIST_GET_GRPC_CLIENT_FAILED, ERROR_DIST_GET_GRPC_CLIENT_FAILED)
		} else {
			resp = client.SendTopicMessage(req)
		}
		if resp == nil {
			xlog.Warn(ERROR_CODE_DIST_GRPC_SERVICE_FAILURE, ERROR_DIST_GRPC_SERVICE_FAILURE)
			producer := s.getProducer(serverId)
			if producer == nil {
				xlog.Warn(ERROR_CODE_DIST_GET_KAFKA_PRODUCER_FAILED, ERROR_DIST_GET_KAFKA_PRODUCER_FAILED)
				return
			}
			_, _, err = producer.EnQueue(req, key)
			if err != nil {
				xlog.Warn(err.Error())
			}
		}
	})
}
