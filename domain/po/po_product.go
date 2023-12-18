package po

import (
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
)

type Product struct {
	entity.GormEntityTs
	Id          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`             // ID
	ProductName string `gorm:"column:product_name;NOT NULL" json:"product_name"`           // 产品名称
	ProductType int    `gorm:"column:product_type;default:0;NOT NULL" json:"product_type"` // 类型 6-积分 3-金币
	Description string `gorm:"column:description;NOT NULL" json:"description"`             // 描述
	Price       int64  `gorm:"column:price;default:0;NOT NULL" json:"price"`               // 价格
	Quantity    int64  `gorm:"column:quantity;default:0;NOT NULL" json:"quantity"`         // 数量 包含赠送
	Give        int64  `gorm:"column:give;default:0;NOT NULL" json:"give"`                 // 赠送
	Image       string `gorm:"column:image;NOT NULL" json:"image"`                         // 图片
}

func (p *Product) Currency() string {
	var (
		currency string
	)
	switch pb_enum.WALLET_TYPE(p.ProductType) {
	case pb_enum.WALLET_TYPE_GOLD_COIN:
		currency = "金币"
	case pb_enum.WALLET_TYPE_POINT:
		currency = "积分"
	}
	return currency
}

func (p *Product) TotalQuantity() int64 {
	return p.Quantity
}
