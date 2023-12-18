package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type RedEnvelopeRecordRepository interface {
	CreateRedEnvelopeRecord(p *po.RedEnvelopeRecord) (err error)
	TxCreateRedEnvelopeRecord(tx *gorm.DB, p *po.RedEnvelopeRecord) (err error)
	TxUpdateRedEnvelopeRecord(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64)
}

type redEnvelopeRecordRepository struct {
}

func NewRedEnvelopeRecordRepository() RedEnvelopeRecordRepository {
	return &redEnvelopeRecordRepository{}
}

func (r *redEnvelopeRecordRepository) CreateRedEnvelopeRecord(p *po.RedEnvelopeRecord) (err error) {
	db := xmysql.GetDB()
	err = db.Create(p).Error
	return
}
func (r *redEnvelopeRecordRepository) TxCreateRedEnvelopeRecord(tx *gorm.DB, p *po.RedEnvelopeRecord) (err error) {
	err = tx.Create(p).Error
	return
}

func (r *redEnvelopeRecordRepository) TxUpdateRedEnvelopeRecord(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64) {
	rowsAffected = tx.Model(po.RedEnvelopeRecord{}).Where(u.Query, u.Args...).Updates(u.Values).RowsAffected
	return
}
