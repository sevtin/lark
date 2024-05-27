package pdo_order

import "lark/pkg/utils"

type OrderStatus struct {
	OrderId     uint64 `json:"order_id" field:"order_id"`
	Uid         uint64 `json:"uid" field:"uid"`
	TradeNo     string `json:"trade_no" field:"trade_no"`
	OrderStatus uint8  `json:"order_status" field:"order_status"`
}

/*SELECT uid,trade_no,order_status FROM orders*/

var (
	field_tag_order_status string
)

func (p *OrderStatus) GetFields() string {
	if field_tag_order_status == "" {
		field_tag_order_status = utils.GetFields(*p)
	}
	return field_tag_order_status
}
