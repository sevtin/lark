package repo

import (
	"gorm.io/gorm"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type RedEnvelopeRepository interface {
	TxCreateRedEnvelope(tx *gorm.DB, p *po.RedEnvelope) (err error)
	TxUpdateRedEnvelope(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64)
	TxRedEnvelopeReturn(tx *gorm.DB, q *entity.MysqlQuery) (rt *pdo.RedEnvelopeReturn, err error)
	GetRedEnvelopeStatus(q *entity.MysqlQuery) (status *pdo.RedEnvelopeStatus, err error)
	GetRedEnvelopeInfo(q *entity.MysqlQuery) (info *pdo.RedEnvelopeInfo, err error)
	GetRemainRedEnvelope(q *entity.MysqlQuery) (info *pdo.RemainRedEnvelopeInfo, err error)
}

type redEnvelopeRepository struct {
}

func NewRedEnvelopeRepository() RedEnvelopeRepository {
	return &redEnvelopeRepository{}
}

func (r *redEnvelopeRepository) TxCreateRedEnvelope(tx *gorm.DB, p *po.RedEnvelope) (err error) {
	err = tx.Create(p).Error
	return
}

func (r *redEnvelopeRepository) TxUpdateRedEnvelope(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64) {
	rowsAffected = tx.Model(po.RedEnvelope{}).Where(u.Query, u.Args...).Updates(u.Values).RowsAffected
	return
}

func (r *redEnvelopeRepository) TxRedEnvelopeReturn(tx *gorm.DB, q *entity.MysqlQuery) (rt *pdo.RedEnvelopeReturn, err error) {
	rt = new(pdo.RedEnvelopeReturn)
	err = tx.Model(po.RedEnvelope{}).Select(rt.GetFields()).Where(q.Query, q.Args...).Find(rt).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *redEnvelopeRepository) GetRedEnvelopeStatus(q *entity.MysqlQuery) (status *pdo.RedEnvelopeStatus, err error) {
	status = new(pdo.RedEnvelopeStatus)
	db := xmysql.GetDB()
	//err = db.Clauses(dbresolver.Write).Model(po.RedEnvelope{}).Where(q.Query, q.Args...).Find(status).Error
	err = db.Model(po.RedEnvelope{}).Where(q.Query, q.Args...).Find(status).Error
	return
}

func (r *redEnvelopeRepository) GetRedEnvelopeInfo(q *entity.MysqlQuery) (info *pdo.RedEnvelopeInfo, err error) {
	info = new(pdo.RedEnvelopeInfo)
	db := xmysql.GetDB()
	err = db.Model(po.RedEnvelope{}).Select(info.GetFields()).Where(q.Query, q.Args...).Find(info).Error
	return
}

func (r *redEnvelopeRepository) GetRemainRedEnvelope(q *entity.MysqlQuery) (info *pdo.RemainRedEnvelopeInfo, err error) {
	info = new(pdo.RemainRedEnvelopeInfo)
	db := xmysql.GetDB()
	err = db.Model(po.RedEnvelope{}).Select(info.GetFields()).Where(q.Query, q.Args...).Find(info).Error
	return
}
