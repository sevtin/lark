package biz_chat_invite

import (
	"github.com/jinzhu/copier"
	chat_client "lark/apps/chat/client"
	user_client "lark/apps/user/client"
	"lark/domain/cache"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_dist"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/proto/pb_user"
)

func ConstructChatInviteNotificationMessage(
	inviteReq *pb_invite.InitiateChatInviteReq,
	invitees []*po.ChatInvite,
	chatCache cache.ChatCache,
	userCache cache.UserCache,
	chatClient chat_client.ChatClient,
	userClient user_client.UserClient) (req *pb_dist.ChatInviteNotificationReq, err error) {
	var (
		userInfo    *pb_user.BasicUserInfo
		chatInfo    *pb_chat.ChatInfo
		userSrvMaps map[int64]int64
	)
	// 1、获取邀请人信息
	userInfo = GetUserInfo(inviteReq.InitiatorUid, userCache, userClient)
	if userInfo.Uid == 0 {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITER_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITER_INFO_FAILED)
		return
	}
	// 2、获取被邀请人serverId
	userSrvMaps, err = GetServerIdList(inviteReq.InviteeUids, userCache, userClient)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITEE_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITEE_INFO_FAILED, err.Error())
		return
	}
	if len(userSrvMaps) == 0 {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_INVITEE_INFO_FAILED, ERROR_CHAT_INVITE_GET_INVITEE_INFO_FAILED)
		return
	}
	if inviteReq.ChatType == pb_enum.CHAT_TYPE_GROUP {
		// 3、获取群信息
		chatInfo = GetGroupChatInfo(inviteReq.ChatId, chatCache, chatClient)
		if chatInfo.ChatId == 0 {
			xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_CHAT_INFO_FAILED, ERROR_CHAT_INVITE_GET_CHAT_INFO_FAILED)
			return
		}
	}
	// 4、构建请求参数
	var (
		invite        *po.ChatInvite
		ok            bool
		initiatorInfo = new(pb_invite.InitiatorInfo)
		serverId      int64
	)
	req = &pb_dist.ChatInviteNotificationReq{
		SenderId:       inviteReq.InitiatorUid,
		SenderPlatform: inviteReq.Platform,
		Notifications:  make([]*pb_dist.ChatInviteNotification, 0),
	}
	copier.Copy(initiatorInfo, userInfo)
	for _, invite = range invitees {
		ci := &pb_msg.ChatInvite{
			InviteId:      invite.InviteId,
			CreatedTs:     invite.CreatedTs,
			InviteeUid:    invite.InviteeUid,
			ChatType:      pb_enum.CHAT_TYPE(invite.ChatType),
			InvitationMsg: invite.InvitationMsg,
			InitiatorInfo: initiatorInfo,
			ChatInfo:      nil,
		}
		if serverId, ok = userSrvMaps[invite.InviteeUid]; ok == false {
			continue
		}
		if pb_enum.CHAT_TYPE(invite.ChatType) == pb_enum.CHAT_TYPE_GROUP {
			ci.ChatInfo = chatInfo
		}
		notification := &pb_dist.ChatInviteNotification{
			ReceiverId:       invite.InviteeUid,
			ReceiverServerId: serverId,
			Invite:           ci,
		}
		req.Notifications = append(req.Notifications, notification)
	}
	return
}

func GetServerIdList(inviteIds []int64, userCache cache.UserCache, userClient user_client.UserClient) (userSrvMaps map[int64]int64, err error) {
	var (
		notUids []int64
		resp    *pb_user.GetServerIdListResp
		server  *pb_user.UserServerId
	)
	userSrvMaps, notUids, _ = userCache.GetUserServerList(inviteIds)
	if len(notUids) == 0 {
		return
	}
	resp = userClient.GetServerIdList(&pb_user.GetServerIdListReq{Uids: notUids})
	if resp == nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GRPC_SERVICE_FAILURE, ERROR_CHAT_INVITE_GRPC_SERVICE_FAILURE)
		err = ERR_CHAT_INVITE_QUERY_DB_FAILED
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		err = ERR_CHAT_INVITE_QUERY_DB_FAILED
		return
	}
	for _, server = range resp.List {
		userSrvMaps[server.Uid] = server.ServerId
	}
	return
}

func GetGroupChatInfo(chatId int64, chatCache cache.ChatCache, chatClient chat_client.ChatClient) (chatInfo *pb_chat.ChatInfo) {
	chatInfo, _ = chatCache.GetGroupChatInfo(chatId)
	if chatInfo.ChatId > 0 {
		return
	}
	var (
		req  = &pb_chat.GetChatInfoReq{ChatId: chatId}
		resp *pb_chat.GetChatInfoResp
	)
	resp = chatClient.GetChatInfo(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GRPC_SERVICE_FAILURE, ERROR_CHAT_INVITE_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	chatInfo = resp.ChatInfo
	return
}

func GetUserInfo(uid int64, userCache cache.UserCache, userClient user_client.UserClient) (userInfo *pb_user.BasicUserInfo) {
	userInfo, _ = userCache.GetBasicUserInfo(uid)
	if userInfo.Uid > 0 {
		return
	}
	var (
		req  = &pb_user.GetBasicUserInfoReq{Uid: uid}
		resp *pb_user.GetBasicUserInfoResp
	)
	resp = userClient.GetBasicUserInfo(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GRPC_SERVICE_FAILURE, ERROR_CHAT_INVITE_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	if resp.UserInfo.Uid == 0 {
		xlog.Warn(ERROR_CODE_CHAT_INVITE_GET_USER_INFO_FAILED, ERROR_CHAT_INVITE_GET_USER_INFO_FAILED)
		return
	}
	userInfo = resp.UserInfo
	return
}
