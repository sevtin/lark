package pdo

import "lark/pkg/utils"

var (
	field_tag_chat_creator string
)

type ChatCreator struct {
	Uid      int64  `json:"uid" field:"uid"`             // 用户ID 系统生成
	Nickname string `json:"nickname" field:"nickname"`   // 昵称
	ServerId int64  `json:"server_id" field:"server_id"` // 分配的ws服务器
	Avatar   string `json:"avatar" field:"avatar"`       // 小图 72*72
}

func (p *ChatCreator) GetFields() string {
	if field_tag_chat_creator == "" {
		field_tag_chat_creator = utils.GetFields(*p)
	}
	return field_tag_chat_creator
}
