package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_obj"
)

func (s *distService) sendChatMessage(buf []byte) (err error) {
	var (
		inbox = new(pb_mq.InboxMessage)
	)
	if err = proto.Unmarshal(buf, inbox); err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_UNMARSHAL_ERR, ERROR_DIST_PROTOCOL_UNMARSHAL_ERR, err.Error())
		// 丢弃无法解析的数据
		err = nil
		return
	}
	s.queues <- struct{}{}
	xants.Submit(func() {
		s.messageDistribution(inbox)
	})
	return
}

func (s *distService) messageDistribution(inbox *pb_mq.InboxMessage) {
	defer func() {
		<-s.queues
	}()
	distMembers := s.getChatMembers(inbox.Msg.ChatId)
	s.sendMessage(inbox, distMembers)
}

func (s *distService) sendMessage(inbox *pb_mq.InboxMessage, distMembers map[int64][]*pb_obj.Int64Array) {
	if inbox.Msg.SrvMsgId == 0 {
		return
	}
	if len(distMembers) == 0 {
		return
	}
	var (
		body []byte
		err  error
	)
	body, err = proto.Marshal(inbox.Msg)
	if err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_MARSHAL_ERR, ERROR_DIST_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	for serverId, _ := range distMembers {
		msgReq := &pb_gw.SendTopicMessageReq{
			Topic:          inbox.Topic,
			SubTopic:       inbox.SubTopic,
			Members:        distMembers[serverId],
			SenderId:       inbox.Msg.SenderId,
			SenderPlatform: inbox.Msg.SenderPlatform,
			Body:           body,
		}
		s.asyncSendMessage(msgReq, serverId, constant.CONST_MSG_KEY_PUSH_ONLINE)
	}
	return
}

func (s *distService) asyncSendMessage(req *pb_gw.SendTopicMessageReq, serverId int64, key string) {
	xants.Submit(func() {
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
