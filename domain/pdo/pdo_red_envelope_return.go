package pdo

import "lark/pkg/utils"

var (
	field_tag_red_envelope_return string
)

type RedEnvelopeReturn struct {
	EnvId        int64 `json:"env_id" field:"env_id"`
	WalletId     int64 `json:"wallet_id" field:"wallet_id"`
	SenderUid    int64 `json:"sender_uid" field:"sender_uid"`
	RemainAmount int64 `json:"remain_amount" field:"remain_amount"`
}

func (p *RedEnvelopeReturn) GetFields() string {
	if field_tag_red_envelope_return == "" {
		field_tag_red_envelope_return = utils.GetFields(*p)
	}
	return field_tag_red_envelope_return
}
