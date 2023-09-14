package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/entity"
)

type PaymentRepository interface {
	TxCreatePayment(tx *gorm.DB, payment *po.Payment) (err error)
	TxUpdatePayment(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64)
}

type paymentRepository struct {
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{}
}

func (o *paymentRepository) TxCreatePayment(tx *gorm.DB, payment *po.Payment) (err error) {
	err = tx.Create(payment).Error
	return
}

func (o *paymentRepository) TxUpdatePayment(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64) {
	rowsAffected = tx.Model(po.Payment{}).Where(u.Query, u.Args...).Updates(u.Values).RowsAffected
	return
}
