package service

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/protocol"
)

func (s *messageService) verifyMessage(req *pb_msg.SendChatMessageReq) (err error) {
	switch req.Msg.MsgType {
	case pb_enum.MSG_TYPE_TEXT:
		if len(req.Msg.Body) == 0 {
			err = ERROR_MESSAGE_BODY_TEXT_EMPTY_ERR
		}
	case pb_enum.MSG_TYPE_IMAGE:
		var (
			body    = new(protocol.Image)
			content = new(pb_msg.Image)
		)
		proto.Unmarshal(req.Msg.Body, content)
		copier.Copy(body, content)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_FILE:
		var (
			body    = new(protocol.File)
			content = new(pb_msg.File)
		)
		proto.Unmarshal(req.Msg.Body, content)
		copier.Copy(body, content)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_AUDIO:
		var (
			body    = new(protocol.Audio)
			content = new(pb_msg.Audio)
		)
		proto.Unmarshal(req.Msg.Body, content)
		copier.Copy(body, content)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_MEDIA:
		var (
			body    = new(protocol.Media)
			content = new(pb_msg.Media)
		)
		proto.Unmarshal(req.Msg.Body, content)
		copier.Copy(body, content)
		err = s.validate.Struct(body)
	case pb_enum.MSG_TYPE_STICKER:
		var (
			body    = new(protocol.Sticker)
			content = new(pb_msg.Sticker)
		)
		proto.Unmarshal(req.Msg.Body, content)
		copier.Copy(body, content)
		err = s.validate.Struct(body)
	}
	return
}
