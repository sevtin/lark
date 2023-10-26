package po

import "lark/pkg/entity"

type Order struct {
	entity.GormEntityTs
	OrderId     int64  `gorm:"column:order_id;primary_key" json:"order_id"`                // order id
	OrderSn     string `gorm:"column:order_sn;NOT NULL" json:"order_sn"`                   // 订单号
	TradeNo     string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                   // 自编唯一交易编号
	Uid         int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`                   // uid
	TimeExpire  int64  `gorm:"column:time_expire;default:0;NOT NULL" json:"time_expire"`   // 绝对超时时间
	Amount      int64  `gorm:"column:amount;default:0;NOT NULL" json:"amount"`             // 订单总金额
	PayType     int    `gorm:"column:pay_type;default:0;NOT NULL" json:"pay_type"`         // 支付方式 1-支付宝 2-微信 3-银联 4-PayPal
	SourceType  int    `gorm:"column:source_type;default:0;NOT NULL" json:"source_type"`   // 订单来源 1-IOS 2-ANDROID 3-MAC 4-WINDOWS 5-WEB
	OrderStatus int    `gorm:"column:order_status;default:0;NOT NULL" json:"order_status"` // 订单状态 0-PENDING 1-PAID 2-CANCELLED 3-REFUNDED
	Integration int64  `gorm:"column:integration;default:0;NOT NULL" json:"integration"`   // 可以获得的积分
	Growth      int64  `gorm:"column:growth;default:0;NOT NULL" json:"growth"`             // 可以获得的成长值
	PaymentTs   int64  `gorm:"column:payment_ts;default:0;NOT NULL" json:"payment_ts"`     // 支付时间
	Subject     string `gorm:"column:subject;NOT NULL" json:"subject"`                     // 订单标题
	Note        string `gorm:"column:note;NOT NULL" json:"note"`                           // 订单备注
	TagId       string `gorm:"column:tag_id;NOT NULL" json:"tag_id"`                       // Tag ID 用于取消
}

type OrderItem struct {
	entity.GormEntityTs
	OrderItemId       int64  `gorm:"column:order_item_id;primary_key" json:"order_item_id"`                  // order item id
	OrderId           int64  `gorm:"column:order_id;default:0;NOT NULL" json:"order_id"`                     // order id
	SpuId             int64  `gorm:"column:spu_id;default:0;NOT NULL" json:"spu_id"`                         // spu id
	SpuName           string `gorm:"column:spu_name;NOT NULL" json:"spu_name"`                               // spu name
	SpuPic            string `gorm:"column:spu_pic;NOT NULL" json:"spu_pic"`                                 // 商品spu图片
	CatId             int64  `gorm:"column:cat_id;default:0;NOT NULL" json:"cat_id"`                         // 商品分类id
	SkuId             int64  `gorm:"column:sku_id;default:0;NOT NULL" json:"sku_id"`                         // sku id
	SkuName           string `gorm:"column:sku_name;NOT NULL" json:"sku_name"`                               // sku name
	SkuPic            string `gorm:"column:sku_pic;NOT NULL" json:"sku_pic"`                                 // 商品sku图片
	SkuPrice          int64  `gorm:"column:sku_price;default:0;NOT NULL" json:"sku_price"`                   // 商品sku价格
	SkuQuantity       int    `gorm:"column:sku_quantity;default:0;NOT NULL" json:"sku_quantity"`             // 商品购买的数量
	SkuAttrs          string `gorm:"column:sku_attrs;NOT NULL" json:"sku_attrs"`                             // 商品销售属性组合（JSON）
	PromotionAmount   int64  `gorm:"column:promotion_amount;default:0;NOT NULL" json:"promotion_amount"`     // 商品促销分解金额
	CouponAmount      int64  `gorm:"column:coupon_amount;default:0;NOT NULL" json:"coupon_amount"`           // 优惠券优惠分解金额
	IntegrationAmount int64  `gorm:"column:integration_amount;default:0;NOT NULL" json:"integration_amount"` // 积分优惠分解金额
	RealAmount        int64  `gorm:"column:real_amount;default:0;NOT NULL" json:"real_amount"`               // 该商品经过优惠后的分解金额
	GiftIntegration   int    `gorm:"column:gift_integration;default:0;NOT NULL" json:"gift_integration"`     // 赠送积分
	GiftGrowth        int    `gorm:"column:gift_growth;default:0;NOT NULL" json:"gift_growth"`               // 赠送成长值
}
