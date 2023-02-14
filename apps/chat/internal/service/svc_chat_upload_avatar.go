package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
)

func (s *chatService) UploadAvatar(ctx context.Context, req *pb_chat.UploadAvatarReq) (resp *pb_chat.UploadAvatarResp, _ error) {
	resp = &pb_chat.UploadAvatarResp{Avatar: &pb_chat.AvatarInfo{}}
	var (
		u   = entity.NewMysqlUpdate()
		err error
	)
	u.Set("avatar_small", req.AvatarSmall)
	u.Set("avatar_medium", req.AvatarMedium)
	u.Set("avatar_large", req.AvatarLarge)
	u.SetFilter("owner_id=?", req.OwnerId)
	u.SetFilter("owner_type=?", pb_enum.AVATAR_OWNER_CHAT_AVATAR)

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		err = s.avatarRepo.TxUpdateAvatar(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED, err.Error())
			return
		}

		u.Reset()
		u.SetFilter("chat_id=?", req.OwnerId)
		u.Set("avatar_key", req.AvatarSmall)
		err = s.chatRepo.TxUpdateChat(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED, err.Error())
			return
		}

		u.Reset()
		u.SetFilter("chat_id=?", req.OwnerId)
		u.Set("chat_avatar_key", req.AvatarSmall)
		err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED)
			xlog.Warn(ERROR_CODE_CHAT_SET_AVATAR_FAILED, ERROR_CHAT_SET_AVATAR_FAILED, err.Error())
			return
		}
		return
	})
	if err != nil {
		return
	}
	err = s.chatCache.DelChatInfo(req.OwnerId)
	if err != nil {
		return
	}
	copier.Copy(resp.Avatar, req)
	return
}
