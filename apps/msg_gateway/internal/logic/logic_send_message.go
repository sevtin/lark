package logic

import (
	"lark/pkg/proto/pb_cm"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_obj"
)

const (
	MAX_NUMBER_OF_PUSHES = 500
)

type SendOnlineMessageHandler func(uid int64, platform int32, message []byte) (result int32)
type SendCloudMessageHandler func(req *pb_cm.CloudMessageReq)

func SendMessage(message *pb_gw.SendMessage, msgBuf []byte, onlineMsgHandler SendOnlineMessageHandler, cloudMsgHandler SendCloudMessageHandler) {
	if len(message.Members) == 0 {
		return
	}
	if len(msgBuf) == 0 {
		return
	}
	var (
		member     *pb_obj.Int64Array
		result     int32
		cmMembers  = make([]*pb_cm.CloudMessageMember, 0)
		uid        int64
		platform   pb_enum.PLATFORM_TYPE
		chatStatus pb_enum.CHAT_STATUS
	)
	// 0:ServerId, 1:Platform, 2:Uid, 3:Status
	for _, member = range message.Members {
		uid = member.GetUid()
		platform = pb_enum.PLATFORM_TYPE(member.GetPlatform())
		chatStatus = pb_enum.CHAT_STATUS(member.GetStatus())

		result = onlineMsgHandler(uid, int32(platform), msgBuf)
		if chatStatus == pb_enum.CHAT_STATUS_MUTE {
			continue
		}
		if uid == message.SenderId {
			continue
		}
		if result > 0 {
			cmMember := &pb_cm.CloudMessageMember{
				Uid:      uid,
				Platform: platform,
			}
			switch platform {
			case pb_enum.PLATFORM_TYPE_IOS, pb_enum.PLATFORM_TYPE_ANDROID:
				if len(cmMembers) <= MAX_NUMBER_OF_PUSHES {
					cmMembers = append(cmMembers, cmMember)
				}
			}
		}
	}
	sendCloudMsg(message, cmMembers, cloudMsgHandler)
}

func sendCloudMsg(message *pb_gw.SendMessage, members []*pb_cm.CloudMessageMember, cloudMsgHandler SendCloudMessageHandler) {
	if len(members) == 0 {
		return
	}
	req := &pb_cm.CloudMessageReq{
		Topic:    message.Topic,
		SubTopic: message.SubTopic,
		Member:   members,
		Body:     message.Body,
	}
	cloudMsgHandler(req)
}
