package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/protocol"
)

func (s *userService) UploadAvatar(ctx context.Context, req *pb_user.UploadAvatarReq) (resp *pb_user.UploadAvatarResp, _ error) {
	resp = &pb_user.UploadAvatarResp{Avatar: &pb_user.AvatarInfo{}}
	var (
		u      = entity.NewMysqlUpdate()
		result *protocol.Result
		err    error
	)

	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()

	u.Set("avatar_small", req.AvatarSmall)
	u.Set("avatar_medium", req.AvatarMedium)
	u.Set("avatar_large", req.AvatarLarge)
	u.SetFilter("owner_id=?", req.OwnerId)
	u.SetFilter("owner_type=?", pb_enum.AVATAR_OWNER_USER_AVATAR)

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		err = s.avatarRepo.TxUpdateAvatar(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
			return
		}

		u.Reset()
		u.SetFilter("uid=?", req.OwnerId)
		u.Set("avatar", req.AvatarSmall)
		err = s.userRepo.TxUpdateUser(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
			return
		}

		u.Reset()
		u.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
		u.SetFilter("uid=?", req.OwnerId)
		u.Set("member_avatar", req.AvatarSmall)
		err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
			return
		}

		result, err = s.updateChatMemberCacheInfo(tx, req.OwnerId)
		if err != nil {
			resp.Set(result.Code, result.Msg)
			return
		}
		return
	})
	if err != nil {
		return
	}

	// 删除缓存
	err = s.userCache.DelUserInfo(req.OwnerId)
	if err != nil {
		resp.Set(ERROR_CODE_USER_UPDATE_USER_CACHE_FAILED, ERROR_USER_UPDATE_USER_CACHE_FAILED)
		return
	}
	copier.Copy(resp.Avatar, req)
	return
}
