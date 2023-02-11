package do

import "lark/pkg/proto/pb_enum"

type ChatMemberInfo struct {
	ChatId   int64               `json:"chat_id"`   // chat ID
	Uid      int64               `json:"uid"`       // 用户ID
	Status   pb_enum.CHAT_STATUS `json:"status"`    // NORMAL:正常模式 MUTE:开启免打扰 BANNED:被禁言
	ServerId int32               `json:"server_id"` // 服务器ID
}

type ChatMemberStatus struct {
	ChatId int64               `json:"chat_id"` // chat ID
	Status pb_enum.CHAT_STATUS `json:"status"`  // NORMAL:正常模式 MUTE:开启免打扰 BANNED:被禁言
}
