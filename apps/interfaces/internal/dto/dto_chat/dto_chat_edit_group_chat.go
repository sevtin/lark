package dto_chat

import "lark/apps/interfaces/internal/dto/dto_kv"

type EditGroupChatReq struct {
	ChatId int64             `json:"chat_id" validate:"required,gt=0"` // chat ID
	Kvs    *dto_kv.KeyValues `json:"kvs" validate:"required"`          // 更新字段
}
