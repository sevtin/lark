package xalipay

import "lark/pkg/utils"

type TradeBody struct {
	TradeNo   string `json:"trade_no"`
	TagId     string `json:"tag_id"`
	ProductID int64  `json:"product_id"`
	Price     int64  `json:"price"`
	Quantity  int64  `json:"quantity"`
	Give      int64  `json:"give"`
}

func (t *TradeBody) ToString() string {
	return utils.ToString(t)
}
