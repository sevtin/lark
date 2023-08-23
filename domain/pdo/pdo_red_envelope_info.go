package pdo

import "lark/pkg/utils"

var (
	field_tag_red_envelope_info string
)

type RedEnvelopeInfo struct {
	EnvId        int64  `json:"env_id" field:"env_id"`
	EnvType      int32  `json:"env_type" field:"env_type"`
	ReceiverType int32  `json:"receiver_type" field:"receiver_type"`
	EnvStatus    int    `json:"env_status" field:"env_status"`
	TradeNo      string `json:"trade_no" field:"trade_no"`
	ChatId       int64  `json:"chat_id" field:"chat_id"`
	SenderUid    int64  `json:"sender_uid" field:"sender_uid"`
	Total        int64  `json:"total" field:"total"`       // total
	Quantity     int32  `json:"quantity" field:"quantity"` // quantity
	//RemainQuantity int    `json:"remain_quantity" field:"remain_quantity"` // 剩余红包数量
	//RemainAmount   int    `json:"remain_amount" field:"remain_amount"`     // 剩余红包金额(分)
	Message   string `json:"message" field:"message"`
	ExpiredTs int64  `json:"expired_ts" field:"expired_ts"`
	Receivers string `json:"receivers" field:"receivers"`
}

func (p *RedEnvelopeInfo) GetFields() string {
	if field_tag_red_envelope_info == "" {
		field_tag_red_envelope_info = utils.GetFields(*p)
	}
	return field_tag_red_envelope_info
}

var (
	field_tag_remain_red_envelope_info string
)

type RemainRedEnvelopeInfo struct {
	EnvId          int64 `json:"env_id" field:"env_id"`
	WalletId       int64 `json:"wallet_id" field:"wallet_id"`             // 红包支出钱包ID
	RemainQuantity int   `json:"remain_quantity" field:"remain_quantity"` // 剩余红包数量
	RemainAmount   int   `json:"remain_amount" field:"remain_amount"`     // 剩余红包金额(分)
}

func (p *RemainRedEnvelopeInfo) GetFields() string {
	if field_tag_remain_red_envelope_info == "" {
		field_tag_remain_red_envelope_info = utils.GetFields(*p)
	}
	return field_tag_remain_red_envelope_info
}
