package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/entity"
)

type OrderRepository interface {
	TxCreateOrder(tx *gorm.DB, order *po.Order) (err error)
	TxCreateOrderItems(tx *gorm.DB, items []*po.OrderItem) (err error)
	TxUpdateOrder(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64)
}

type orderRepository struct {
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (o *orderRepository) TxCreateOrder(tx *gorm.DB, order *po.Order) (err error) {
	err = tx.Create(order).Error
	return
}

func (o *orderRepository) TxCreateOrderItems(tx *gorm.DB, items []*po.OrderItem) (err error) {
	err = tx.Create(items).Error
	return
}

func (o *orderRepository) TxUpdateOrder(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64) {
	rowsAffected = tx.Model(po.Order{}).Where(u.Query, u.Args...).Updates(u.Values).RowsAffected
	return
}
