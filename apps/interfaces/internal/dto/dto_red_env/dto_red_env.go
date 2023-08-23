package dto_red_env

type GiveRedEnvelopeReq struct {
	EnvType        int     `json:"env_type"`        // 红包类型 1-均分红包 2-碰运气红包
	ReceiverType   int     `json:"receiver_type"`   // 接收者类型 1-私聊对方 2-群聊所有人 3-群聊指定人
	ChatId         int64   `json:"chat_id"`         // 聊天ID
	SenderUid      int64   `json:"sender_uid"`      // 发红包用户ID
	Total          int64   `json:"total"`           // 红包总金额(分)
	Quantity       int32   `json:"quantity"`        // 红包数量
	Message        string  `json:"message"`         // 祝福语
	ReceiverUids   []int64 `json:"receiver_uids"`   // 接收者ID
	SenderPlatform int32   `json:"sender_platform"` // 发红包平台
}
