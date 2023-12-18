package pdo

import "lark/pkg/utils"

type WithdrawVerify struct {
	WalletId    int64  `json:"wallet_id" field:"wallet_id"`
	WalletType  int32  `json:"wallet_type" field:"wallet_type"`
	Uid         int64  `json:"uid" field:"uid"`
	PayPassword string `json:"pay_password" field:"pay_password"`
	Balance     int64  `json:"balance" field:"balance"`
}

/*
SELECT wallet_id,wallet_type,uid,pay_password,balance
FROM wallets
*/

var (
	field_tag_withdraw_verify string
)

func (p *WithdrawVerify) GetFields() string {
	if field_tag_withdraw_verify == "" {
		field_tag_withdraw_verify = utils.GetFields(*p)
	}
	return field_tag_withdraw_verify
}
