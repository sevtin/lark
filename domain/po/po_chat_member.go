package po

import "lark/pkg/entity"

type ChatMember struct {
	entity.GormEntityTs
	ChatId       int64  `gorm:"column:chat_id;primary_key" json:"chat_id"`                // chat ID
	Uid          int64  `gorm:"column:uid;primary_key;NOT NULL" json:"uid"`               // 用户ID
	ChatType     int    `gorm:"column:chat_type;NOT NULL" json:"chat_type"`               // chat type 1:私聊/2:群聊
	ChatName     string `gorm:"column:chat_name;NOT NULL" json:"chat_name"`               // 名称
	Remark       string `gorm:"column:remark;NOT NULL" json:"remark"`                     // 备注
	OwnerId      int64  `gorm:"column:owner_id;default:0;NOT NULL" json:"owner_id"`       // 归属人ID
	RoleId       int    `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"`         // 角色ID
	Alias        string `gorm:"column:alias;NOT NULL" json:"alias"`                       // 别名
	MemberAvatar string `gorm:"column:member_avatar;NOT NULL" json:"member_avatar"`       // member头像 72*72
	ChatAvatar   string `gorm:"column:chat_avatar;NOT NULL" json:"chat_avatar"`           // chat头像 72*72
	Sync         int    `gorm:"column:sync;default:0;NOT NULL" json:"sync"`               // 是否同步用户信息 0:同步 1:不同步
	Status       int    `gorm:"column:status;default:0;NOT NULL" json:"status"`           // NORMAL:正常模式 MUTE:开启免打扰 BANNED:被禁言
	JoinSource   int    `gorm:"column:join_source;default:0;NOT NULL" json:"join_source"` // 加入源
	ReadSeq      int    `gorm:"column:read_seq;default:0;NOT NULL" json:"read_seq"`       // 已读最后SEQ ID
	Slot         int    `gorm:"column:slot;default:0;NOT NULL" json:"slot"`               // 槽位
}
