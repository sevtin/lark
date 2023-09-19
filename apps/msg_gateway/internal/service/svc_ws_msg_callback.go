package service

import (
	"google.golang.org/protobuf/proto"
	"lark/apps/msg_gateway/internal/server/websocket/ws"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *wsService) replyMessage(message *ws.Message, isSrv bool, code int32, msg string) {
	var (
		resp *pb_msg.MessageResp
		buf  []byte
		err  error
	)
	resp = &pb_msg.MessageResp{
		Code:  code,
		Msg:   msg,
		MsgId: message.Packet.MsgId,
		IsSrv: isSrv,
	}
	buf, err = proto.Marshal(resp)
	if err != nil {
		xlog.Warn(ERROR_CODE_WS_PROTOCOL_MARSHAL_ERR, ERROR_WS_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	buf, err = utils.Encode(int32(message.Packet.Topic), int32(message.Packet.SubTopic), int32(pb_enum.MESSAGE_TYPE_RESP), buf)
	if err != nil {
		xlog.Warn(ERROR_CODE_WS_ASSEMBLY_PROTOCOL_ERR, ERROR_WS_ASSEMBLY_PROTOCOL_ERR, err.Error())
		return
	}
	message.Hub.SendMessage(message.Uid, message.Platform, buf)
}

func (s *wsService) MessageCallback(msg *ws.Message) {
	switch msg.Packet.Topic {
	case pb_enum.TOPIC_CHAT:
		s.chatSubtopic(msg)
	case pb_enum.TOPIC_READ_RECEIPT:
		s.readReceipt(msg)
	default:
		s.replyMessage(msg,
			false,
			ERROR_CODE_WS_TOPIC_ID_ERR,
			ERROR_WS_TOPIC_ID_ERR)
		xlog.Warn(ERROR_CODE_WS_TOPIC_ID_ERR, ERROR_WS_TOPIC_ID_ERR, msg.Packet.Topic)
	}
}

func (s *wsService) chatSubtopic(msg *ws.Message) {
	switch msg.Packet.SubTopic {
	case pb_enum.SUB_TOPIC_CHAT_MSG:
		s.sendChatMessage(msg)
	default:
		s.replyMessage(msg,
			false,
			ERROR_CODE_WS_TOPIC_ID_ERR,
			ERROR_WS_TOPIC_ID_ERR)
		xlog.Warn(ERROR_CODE_WS_TOPIC_ID_ERR, ERROR_WS_TOPIC_ID_ERR, msg.Packet.Topic)
	}
}

func (s *wsService) sendChatMessage(msg *ws.Message) {
	var (
		req  *pb_msg.SendChatMessageReq
		resp *pb_msg.SendChatMessageResp
		err  error
	)
	req = &pb_msg.SendChatMessageReq{
		Topic:    msg.Packet.Topic,
		SubTopic: msg.Packet.SubTopic,
		Msg:      new(pb_msg.CliChatMessage),
	}
	err = proto.Unmarshal(msg.Packet.Data, req.Msg)
	if err != nil {
		s.replyMessage(msg,
			false,
			ERROR_CODE_WS_PROTOCOL_UNMARSHAL_ERR,
			ERROR_WS_PROTOCOL_UNMARSHAL_ERR)
		xlog.Warn(ERROR_CODE_WS_PROTOCOL_UNMARSHAL_ERR, ERROR_WS_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	req.Msg.SenderId = msg.Uid
	req.Msg.SenderPlatform = pb_enum.PLATFORM_TYPE(msg.Platform)
	resp = s.msgClient.SendChatMessage(req)
	if resp == nil {
		s.replyMessage(msg,
			false,
			ERROR_CODE_WS_GRPC_SERVICE_FAILURE,
			ERROR_WS_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_WS_GRPC_SERVICE_FAILURE, ERROR_WS_GRPC_SERVICE_FAILURE)
		return
	}
	s.replyMessage(msg,
		false,
		resp.Code,
		resp.Msg)
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
}

func (s *wsService) readReceipt(msg *ws.Message) {
	var (
		receipt = new(pb_msg.ReadReceipt)
		err     error
	)
	err = proto.Unmarshal(msg.Packet.Data, receipt)
	if err != nil {
		return
	}
	receipt.Uid = msg.Uid
	s.producer.EnQueue(receipt, constant.CONST_MSG_KEY_READ_RECEIPT)
}
