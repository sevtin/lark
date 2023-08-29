package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
)

func (s *messageService) verifyMessage(req *pb_msg.SendChatMessageReq) (assocId int64,err error) {
	switch req.Msg.MsgType {
	case pb_enum.MSG_TYPE_TEXT:
		if len(req.Msg.Body) > MAX_MESSAGE_LENGTH {
			err = ERROR_MESSAGE_OUT_MAX_RANGE
			return
		}
		if len(req.Msg.Body) == 0 {
			err = ERROR_MESSAGE_BODY_TEXT_EMPTY_ERR
		}
	case pb_enum.MSG_TYPE_IMAGE:
		var (
			body = new(pb_msg.Image)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_FILE:
		var (
			body = new(pb_msg.File)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_AUDIO:
		var (
			body = new(pb_msg.Audio)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_MEDIA:
		var (
			body = new(pb_msg.Media)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_STICKER:
		var (
			body = new(pb_msg.Sticker)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_GIVE_RED_ENV:
		var (
			body = new(pb_msg.GiveRedEnvelope)
		)
		proto.Unmarshal(req.Msg.Body, body)
		assocId = body.EnvId
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_RECEIVE_RED_ENV:
		var (
			body = new(pb_msg.ReceiveRedEnvelope)
		)
		proto.Unmarshal(req.Msg.Body, body)
		err = s.validate.Struct(body)
	}
	return
}
