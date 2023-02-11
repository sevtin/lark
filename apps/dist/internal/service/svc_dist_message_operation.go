package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/proto/pb_obj"
)

func (s *distService) messageOperation(buf []byte) (err error) {
	var (
		req           = new(pb_msg.MessageOperationReq)
		serverMembers map[int64][]*pb_obj.Int64Array
	)
	if err = proto.Unmarshal(buf, req); err != nil {
		xlog.Warn(ERROR_CODE_DIST_PROTOCOL_UNMARSHAL_ERR, ERROR_DIST_PROTOCOL_UNMARSHAL_ERR, err.Error())
		// 丢弃无法解析的数据
		err = nil
		return
	}
	serverMembers = s.getChatMembers(req.Operation.ChatId)
	s.sendMessageOperation(req, serverMembers)
	return
}

func (s *distService) sendMessageOperation(opnReq *pb_msg.MessageOperationReq, distMembers map[int64][]*pb_obj.Int64Array) {
	if len(distMembers) == 0 {
		return
	}
	var (
		serverId int64
		body     []byte
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
		s.asyncSendMessage(msgReq, serverId, constant.CONST_MSG_KEY_PUSH_OPERATION)
	}
	return
}
