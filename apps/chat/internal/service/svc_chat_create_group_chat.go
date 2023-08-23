package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"lark/business/biz_chat_invite"
	"lark/domain/do"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_dist"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_invite"
)

func (s *chatService) CreateGroupChat(ctx context.Context, req *pb_chat.CreateGroupChatReq) (resp *pb_chat.CreateGroupChatResp, _ error) {
	resp = &pb_chat.CreateGroupChatResp{}
	var (
		creator = new(pdo.ChatCreator)
		q       = entity.NewMysqlQuery()
		chat    *po.Chat
		err     error
	)
	var (
		avatar        *po.Avatar
		member        *po.ChatMember
		invitationMsg string
		uid           int64
		invite        *po.ChatInvite
		inviteList    = make([]*po.ChatInvite, 0)
	)
	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()

	// 1 获取创建者信息
	q.Fields = creator.GetFields()
	q.SetFilter("uid=?", req.CreatorUid)
	err = s.userRepo.QueryUser(q, creator)
	if err != nil || creator.Uid == 0 {
		resp.Set(ERROR_CODE_CHAT_QUERY_DB_FAILED, ERROR_CHAT_QUERY_DB_FAILED)
		return
	}

	// 2 构建chat模型
	chat = &po.Chat{
		CreatorUid: req.CreatorUid,
		ChatType:   int(pb_enum.CHAT_TYPE_GROUP),
		Avatar:     constant.CONST_AVATAR_SMALL,
		Name:       req.Name,
		About:      req.About,
	}

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		// 3 chat入库
		err = s.chatRepo.TxCreate(tx, chat)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
			return
		}

		// 4 creator入群/入库
		member = &po.ChatMember{
			ChatId:       chat.ChatId,
			ChatType:     chat.ChatType,
			ChatName:     chat.Name,
			Uid:          creator.Uid,
			RoleId:       int(pb_enum.CHAT_GROUP_ROLE_MASTER),
			Alias:        creator.Nickname,
			MemberAvatar: creator.Avatar,
			ChatAvatar:   chat.Avatar,
			Sync:         constant.SYNCHRONIZE_USER_INFO,
		}
		err = s.chatMemberRepo.TxCreate(tx, member)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
			return
		}

		// 5 设置群头像
		avatar = &po.Avatar{
			OwnerId:      chat.ChatId,
			OwnerType:    int(pb_enum.AVATAR_OWNER_CHAT_AVATAR),
			AvatarSmall:  constant.CONST_AVATAR_SMALL,
			AvatarMedium: constant.CONST_AVATAR_MEDIUM,
			AvatarLarge:  constant.CONST_AVATAR_LARGE,
		}
		err = s.avatarRepo.TxCreate(tx, avatar)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
			return
		}
		// 6 构建邀请信息
		invitationMsg = creator.Nickname + CONST_CHAT_INVITE_TITLE_CONJUNCTION + chat.Name
		for _, uid = range req.UidList {
			if uid == req.CreatorUid {
				continue
			}
			invite = &po.ChatInvite{
				InviteId:      xsnowflake.NewSnowflakeID(),
				ChatId:        chat.ChatId,
				ChatType:      chat.ChatType,
				InitiatorUid:  req.CreatorUid,
				InviteeUid:    uid,
				InvitationMsg: invitationMsg,
			}
			inviteList = append(inviteList, invite)
		}
		if len(inviteList) == 0 {
			return
		}
		// 7 邀请信息入库
		err = s.chatInviteRepo.TxCreateChatInvites(tx, inviteList)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_INSERT_VALUE_FAILED, ERROR_CHAT_INSERT_VALUE_FAILED)
			return
		}
		return
	})
	if err != nil {
		return
	}

	xants.Submit(func() {
		var (
			err error
		)
		// 8 缓存成员hash
		err = s.chatMemberCache.HSetNXChatMember(member.ChatId, member.Uid, fmt.Sprintf("%d,%d", creator.ServerId, member.Status))
		if err != nil {
			xlog.Warn(err.Error())
			var (
				kfv = do.KeyFieldValue{member.ChatId, member.Uid, fmt.Sprintf("%d,%d", creator.ServerId, member.Status)}
			)
			_, _, err = s.cacheProducer.Push(kfv, constant.CONST_MSG_KEY_CACHE_CREATE_GROUP_CHAT)
			if err != nil {
				xlog.Warn(err.Error())
			}
		}
		// 9 邀请推送
		inviteReq := &pb_invite.InitiateChatInviteReq{
			ChatId:        chat.ChatId,
			ChatType:      pb_enum.CHAT_TYPE(chat.ChatType),
			InitiatorUid:  req.CreatorUid,
			InviteeUids:   req.UidList,
			InvitationMsg: invitationMsg,
			Platform:      0,
		}
		s.sendChatInviteNotificationMessage(inviteReq, inviteList)
	})
	return
}

func (s *chatService) sendChatInviteNotificationMessage(inviteReq *pb_invite.InitiateChatInviteReq, invitees []*po.ChatInvite) {
	var (
		req  *pb_dist.ChatInviteNotificationReq
		resp *pb_dist.ChatInviteNotificationResp
		err  error
	)
	req, err = biz_chat_invite.ConstructChatInviteNotificationMessage(
		inviteReq,
		invitees,
		s.chatCache,
		s.userCache,
		s.chatRepo,
		s.userRepo)
	if err != nil {
		return
	}
	if req == nil {
		return
	}
	// TODO: rpc error: code = ResourceExhausted desc = grpc: received message larger than max (25156 vs. 4096)
	resp = s.distClient.ChatInviteNotification(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_CHAT_GRPC_SERVICE_FAILURE, ERROR_CHAT_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
	}
}
