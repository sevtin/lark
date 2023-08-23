package pdo

import "lark/pkg/utils"

var (
	field_tag_account_info string
)

type AccountInfo struct {
	WalletId     int64 `json:"wallet_id" field:"wallet_id"`
	WalletType   int32 `json:"wallet_type" field:"wallet_type"`
	Uid          int64 `json:"uid" field:"uid"`
	Balance      int64 `json:"balance" field:"balance"`
	FrozenAmount int64 `json:"frozen_amount" field:"frozen_amount"`
	Status       int32 `json:"status" field:"status"`
}

func (p *AccountInfo) GetFields() string {
	if field_tag_account_info == "" {
		field_tag_account_info = utils.GetFields(*p)
	}
	return field_tag_account_info
}
