package dto_user

import "lark/apps/interfaces/internal/dto/dto_kv"

type EditUserInfoReq struct {
	Kvs *dto_kv.KeyValues `json:"kvs" binding:"required"` // 更新字段
}
