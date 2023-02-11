package gateway

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/utils"
)

func (s *gatewayServer) SendTopicMessage(ctx context.Context, req *pb_gw.SendTopicMessageReq) (resp *pb_gw.SendTopicMessageResp, _ error) {
	resp = new(pb_gw.SendTopicMessageResp)
	switch req.Topic {
	case pb_enum.TOPIC_CHAT:
		s.topicMessageHandler(req, resp)
	case pb_enum.TOPIC_CHAT_INVITE:
		s.topicInvite(req, resp)
	}
	return
}

func (s *gatewayServer) topicMessageHandler(req *pb_gw.SendTopicMessageReq, resp *pb_gw.SendTopicMessageResp) {
	switch req.SubTopic {
	case pb_enum.SUB_TOPIC_CHAT_MSG,
		pb_enum.SUB_TOPIC_CHAT_JOINED_GROUP_CHAT,
		pb_enum.SUB_TOPIC_CHAT_QUIT_GROUP_CHAT,
		pb_enum.SUB_TOPIC_CHAT_REMOVE_CHAT_MEMBER:
		s.sendChatMessage(req, resp)
	case pb_enum.SUB_TOPIC_CHAT_OPERATION:
		s.sendOperationMessage(req, resp)
	}
	return
}

func (s *gatewayServer) topicInvite(req *pb_gw.SendTopicMessageReq, resp *pb_gw.SendTopicMessageResp) {
	switch req.SubTopic {
	case pb_enum.SUB_TOPIC_CHAT_INVITE_REQUEST:
		s.sendInviteMessage(req, resp)
	}
}

func AssemblyMessage(req *pb_gw.SendTopicMessageReq, senderId int64, senderPlatform pb_enum.PLATFORM_TYPE, msgType pb_enum.MESSAGE_TYPE, resp *pb_gw.SendTopicMessageResp) (message *pb_gw.SendMessage, buf []byte, err error) {
	message = &pb_gw.SendMessage{
		Topic:          req.Topic,
		SubTopic:       req.SubTopic,
		Members:        req.Members,
		SenderId:       senderId,
		SenderPlatform: senderPlatform,
		Body:           req.Body,
	}
	buf, err = Encode(req.Topic, req.SubTopic, msgType, req.Body, resp)
	return
}

func Encode(topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, msgType pb_enum.MESSAGE_TYPE, body []byte, resp *pb_gw.SendTopicMessageResp) (buf []byte, err error) {
	buf, err = utils.Encode(int32(topic), int32(subTopic), int32(msgType), body)
	if err != nil {
		resp.Set(ERROR_CODE_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, ERROR_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR)
		xlog.Warn(ERROR_CODE_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, ERROR_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, err)
	}
	return
}
