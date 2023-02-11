package service

import (
	"context"
	"google.golang.org/protobuf/proto"
	"lark/apps/dist/internal/logic"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_dist"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_obj"
)

func (s *distService) ChatInviteNotification(ctx context.Context, req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp, err error) {
	var (
		index        int
		notification *pb_dist.ChatInviteNotification
		serverId     int64
	)
	for index, _ = range req.Notifications {
		notification = req.Notifications[index]
		var (
			distMembers map[int64][]*pb_obj.Int64Array
			body        []byte
		)
		body, err = proto.Marshal(notification.Invite)
		if err != nil {
			xlog.Warn(ERROR_CODE_DIST_PROTOCOL_MARSHAL_ERR, ERROR_DIST_PROTOCOL_MARSHAL_ERR, err.Error())
			return
		}
		distMembers = logic.GetDistMembers(notification.ReceiverServerId, notification.ReceiverId, 0)
		for serverId, _ = range distMembers {
			msgReq := &pb_gw.SendTopicMessageReq{
				Topic:          pb_enum.TOPIC_CHAT_INVITE,
				SubTopic:       pb_enum.SUB_TOPIC_CHAT_INVITE_REQUEST,
				Members:        distMembers[serverId],
				SenderId:       notification.Invite.InitiatorInfo.Uid,
				SenderPlatform: 0,
				Body:           body,
			}
			s.asyncSendMessage(msgReq, serverId, constant.CONST_MSG_KEY_CHAT_INVITE)
		}
	}
	return
}
