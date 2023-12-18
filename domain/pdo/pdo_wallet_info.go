package pdo

import "lark/pkg/utils"

type WalletInfo struct {
	WalletId   int64 `json:"wallet_id" field:"wallet_id"`
	WalletType int32 `json:"wallet_type" field:"wallet_type"`
	Uid        int64 `json:"uid" field:"uid"`
	Balance    int64 `json:"balance" field:"balance"`
	Status     int32 `json:"status" field:"status"`
}

/*
SELECT wallet_id,wallet_type,uid,balance,status
FROM wallets
*/

var (
	field_tag_wallet_info string
)

func (p *WalletInfo) GetFields() string {
	if field_tag_wallet_info == "" {
		field_tag_wallet_info = utils.GetFields(*p)
	}
	return field_tag_wallet_info
}
