package po

import "lark/pkg/entity"

type ChatInvite struct {
	entity.GormEntityTs
	InviteId      int64  `gorm:"column:invite_id;primary_key" json:"invite_id"`                // invite ID
	ChatId        int64  `gorm:"column:chat_id;NOT NULL" json:"chat_id"`                       // Chat ID
	ChatType      int    `gorm:"column:chat_type;NOT NULL" json:"chat_type"`                   // 1:私聊/2:群聊
	InitiatorUid  int64  `gorm:"column:initiator_uid;NOT NULL" json:"initiator_uid"`           // 发起人 UID
	InviteeUid    int64  `gorm:"column:invitee_uid;NOT NULL" json:"invitee_uid"`               // 被邀请人 UID
	InvitationMsg string `gorm:"column:invitation_msg;NOT NULL" json:"invitation_msg"`         // 邀请消息
	HandlerUid    int64  `gorm:"column:handler_uid;default:0;NOT NULL" json:"handler_uid"`     // 处理人 UID
	HandleResult  int    `gorm:"column:handle_result;default:0;NOT NULL" json:"handle_result"` // 结果
	HandleMsg     string `gorm:"column:handle_msg" json:"handle_msg"`                          // 处理消息
	HandledTs     int64  `gorm:"column:handled_ts;default:0;NOT NULL" json:"handled_ts"`       // 处理时间
}
