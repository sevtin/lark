package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
)

func (s *chatService) DeleteContact(ctx context.Context, req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp, _ error) {
	resp = new(pb_chat.DeleteContactResp)
	var (
		u   = entity.NewMysqlUpdate()
		err error
	)
	u.SetFilter("chat_id=?", req.ChatId)
	u.SetFilter("owner_id=?", req.Uid)
	u.SetFilter("deleted_ts=?", 0)
	u.Set("status", int32(pb_enum.CHAT_STATUS_DELETED))
	u.Set("deleted_ts", utils.NowUnix())
	_, err = s.removeChatMember(u, req.ChatId, []int64{req.Uid, req.ContactId}, pb_enum.CHAT_TYPE_PRIVATE)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_CHAT_UPDATE_VALUE_FAILED, ERROR_CHAT_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	return
}
