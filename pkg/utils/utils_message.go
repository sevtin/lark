package utils

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
)

func PrivateChatSessionKey(chatId int64, uid1 int64, uid2 int64) string {
	var (
		sessionKey string
	)
	if uid1 > uid2 {
		sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid1) + "-" + Int64ToStr(uid2)
	} else {
		sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid2) + "-" + Int64ToStr(uid1)
	}
	return MD5(sessionKey)
}

func GroupChatSessionKey(chatId int64, uid int64) string {
	var (
		sessionKey string
	)
	sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid)
	return MD5(sessionKey)
}

func MsgBodyToStr(msgType pb_enum.MSG_TYPE, buf []byte) (str string) {
	switch msgType {
	case pb_enum.MSG_TYPE_TEXT:
		str = Bytes2Str(buf)
	case pb_enum.MSG_TYPE_IMAGE:
		var (
			content = new(pb_msg.Image)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_FILE:
		var (
			content = new(pb_msg.File)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_AUDIO:
		var (
			content = new(pb_msg.Audio)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_MEDIA:
		var (
			content = new(pb_msg.Media)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_STICKER:
		var (
			content = new(pb_msg.Sticker)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_JOINED_GROUP_CHAT:
		var (
			content = new(pb_msg.JoinedGroupChatMessage)
		)
		proto.Unmarshal(buf, content)
		str = ToString(content)
	case pb_enum.MSG_TYPE_ACCEPTED_CHAT_INVITE, pb_enum.MSG_TYPE_CHAT_INVITE_MSG, pb_enum.MSG_TYPE_QUIT_GROUP_CHAT, pb_enum.MSG_TYPE_REMOVE_CHAT_MEMBER:
		str = Bytes2Str(buf)
	default:
		str = Bytes2Str(buf)
	}
	return
}
