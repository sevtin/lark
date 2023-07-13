package po

import "lark/pkg/entity"

type Chat struct {
	entity.GormEntityTs
	ChatId     int64  `gorm:"column:chat_id;primary_key" json:"chat_id"`      // chat ID
	CreatorUid int64  `gorm:"column:creator_uid;NOT NULL" json:"creator_uid"` // 创建者 uid
	ChatType   int    `gorm:"column:chat_type;NOT NULL" json:"chat_type"`     // chat type 1:私聊/2:群聊
	Avatar     string `gorm:"column:avatar;NOT NULL" json:"avatar"`           // 小图 72*72
	Name       string `gorm:"column:name" json:"name"`                        // chat标题
	About      string `gorm:"column:about" json:"about"`                      // 关于
}
