package po

import "lark/pkg/entity"

type ChatMember struct {
	entity.GormEntityTs
	ChatId          int64  `gorm:"column:chat_id;primary_key" json:"chat_id"`                  // chat ID
	Uid             int64  `gorm:"column:uid;primary_key;NOT NULL" json:"uid"`                 // 用户ID
	ChatType        int    `gorm:"column:chat_type;NOT NULL" json:"chat_type"`                 // chat type 1:私聊/2:群聊
	ChatName        string `gorm:"column:chat_name;NOT NULL" json:"chat_name"`                 // 名称
	Remark          string `gorm:"column:remark;NOT NULL" json:"remark"`                       // 备注
	OwnerId         int64  `gorm:"column:owner_id;default:0;NOT NULL" json:"owner_id"`         // 归属人ID
	RoleId          int    `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"`           // 角色ID
	Alias           string `gorm:"column:alias;NOT NULL" json:"alias"`                         // 别名
	MemberAvatarKey string `gorm:"column:member_avatar_key;NOT NULL" json:"member_avatar_key"` // member头像 72*72
	ChatAvatarKey   string `gorm:"column:chat_avatar_key;NOT NULL" json:"chat_avatar_key"`     // chat头像 72*72
	Sync            int    `gorm:"column:sync;default:0;NOT NULL" json:"sync"`                 // 是否同步用户信息 0:同步 1:不同步
	Status          int    `gorm:"column:status;default:0;NOT NULL" json:"status"`             // NORMAL:正常模式 MUTE:开启免打扰 BANNED:被禁言
	ServerId        int64  `gorm:"column:server_id;default:0;NOT NULL" json:"server_id"`       // 服务器ID
	JoinSource      int    `gorm:"column:join_source;default:0;NOT NULL" json:"join_source"`   // 加入源
}
