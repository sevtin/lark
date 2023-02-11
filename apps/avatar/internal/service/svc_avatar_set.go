package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_avatar"
	"lark/pkg/proto/pb_enum"
)

func (s *avatarService) SetAvatar(ctx context.Context, req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp, _ error) {
	resp = &pb_avatar.SetAvatarResp{Avatar: &pb_avatar.AvatarInfo{}}
	var (
		u   = entity.NewMysqlUpdate()
		err error
	)
	u.Set("avatar_small", req.AvatarSmall)
	u.Set("avatar_medium", req.AvatarMedium)
	u.Set("avatar_large", req.AvatarLarge)
	u.SetFilter("owner_id=?", req.OwnerId)
	u.SetFilter("owner_type=?", int32(req.OwnerType))

	tx := xmysql.GetTX()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = s.avatarRepo.TxUpdateAvatar(tx, u)
	if err != nil {
		resp.Set(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED, err.Error())
		return
	}
	u.Reset()
	switch req.OwnerType {
	case pb_enum.AVATAR_OWNER_USER_AVATAR:
		u.SetFilter("uid=?", req.OwnerId)
		u.Set("avatar_key", req.AvatarSmall)
		s.userRepo.UpdateUser(u)
		if err != nil {
			resp.Set(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED)
			xlog.Warn(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED, err.Error())
			return
		}

		u.Reset()
		u.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
		u.SetFilter("uid=?", req.OwnerId)
		u.Set("member_avatar_key", req.AvatarSmall)
	case pb_enum.AVATAR_OWNER_CHAT_AVATAR:
		u.SetFilter("chat_id=?", req.OwnerId)
		u.Set("avatar_key", req.AvatarSmall)
		err = s.chatRepo.TxUpdateChat(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED)
			xlog.Warn(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED, err.Error())
			return
		}

		u.Reset()
		u.SetFilter("chat_id=?", req.OwnerId)
		u.Set("chat_avatar_key", req.AvatarSmall)
	}

	err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
	if err != nil {
		resp.Set(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_AVATAR_SET_AVATAR_FAILED, ERROR_AVATAR_SET_AVATAR_FAILED, err.Error())
		return
	}
	copier.Copy(resp.Avatar, req)
	return
}
