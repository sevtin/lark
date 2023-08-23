package pdo

import "lark/pkg/utils"

var (
	field_tag_account_balance string
)

type AccountBalance struct {
	WalletId int64 `json:"wallet_id" field:"wallet_id"`
	Balance  int64 `json:"balance" field:"balance"`
}

func (p *AccountBalance) GetFields() string {
	if field_tag_account_balance == "" {
		field_tag_account_balance = utils.GetFields(*p)
	}
	return field_tag_account_balance
}
