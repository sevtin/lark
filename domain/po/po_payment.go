package po

import "lark/pkg/entity"

type Payment struct {
	entity.GormEntityTs
	PayId         int64  `gorm:"column:pay_id;primary_key" json:"pay_id"`                      // pay id
	Uid           int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`                     // uid
	SellerId      string `gorm:"column:seller_id;NOT NULL" json:"seller_id"`                   // 收款账号对应的第三方支付唯一用户号
	BuyerId       string `gorm:"column:buyer_id;NOT NULL" json:"buyer_id"`                     // 支付人所在支付平台ID
	OrderId       int64  `gorm:"column:order_id;default:0;NOT NULL" json:"order_id"`           // 商户原始订单号
	TradeNo       string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                     // 系统内部唯一订单号，只能是数字、大小写字母_-*
	PayStatus     int    `gorm:"column:pay_status;default:0;NOT NULL" json:"pay_status"`       // 支付状态 0-待支付 1-已支付/已完成 2-已取消 3-失败
	Currency      string `gorm:"column:currency;NOT NULL" json:"currency"`                     // 币种
	Subject       string `gorm:"column:subject;NOT NULL" json:"subject"`                       // 订单标题
	Summary       string `gorm:"column:summary;NOT NULL" json:"summary"`                       // 摘要
	ThTradeNo     string `gorm:"column:th_trade_no;NOT NULL" json:"th_trade_no"`               // 该交易在第三方支付系统中的交易流水号
	PayType       int    `gorm:"column:pay_type;default:0;NOT NULL" json:"pay_type"`           // 支付方式 1-支付宝 2-微信 3-银联 4-Paypal
	ReturnContent string `gorm:"column:return_content;NOT NULL" json:"return_content"`         // return内容
	NotifyContent string `gorm:"column:notify_content;NOT NULL" json:"notify_content"`         // notify内容
	ResultContent string `gorm:"column:result_content;NOT NULL" json:"result_content"`         // 结果内容
	ReturnTs      int64  `gorm:"column:return_ts;default:0;NOT NULL" json:"return_ts"`         // return时间
	NotifyTs      int64  `gorm:"column:notify_ts;default:0;NOT NULL" json:"notify_ts"`         // notify时间
	TotalAmount   int64  `gorm:"column:total_amount;default:0;NOT NULL" json:"total_amount"`   // 订单金额
	ActualAmount  int64  `gorm:"column:actual_amount;default:0;NOT NULL" json:"actual_amount"` // 实际入账金额
	PayTs         int64  `gorm:"column:pay_ts;default:0;NOT NULL" json:"pay_ts"`               // 支付时间
	TagId         string `gorm:"column:tag_id;NOT NULL" json:"tag_id"`                         // Tag ID 用于取消
	SaleId        string `gorm:"column:sale_id;NOT NULL" json:"sale_id"`                       // Sale ID 用于退款
}
