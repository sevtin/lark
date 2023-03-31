package repo

import (
	"gorm.io/gorm"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type ChatInviteRepository interface {
	CreateChatInvites(invites []*po.ChatInvite) (err error)
	TxCreateChatInvites(tx *gorm.DB, list []*po.ChatInvite) (err error)
	TxUpdateChatInvite(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64)
	TxChatInvite(tx *gorm.DB, w *entity.MysqlWhere) (invite *po.ChatInvite, err error)
	ChatInvite(w *entity.MysqlWhere) (invite *po.ChatInvite, err error)
	ChatInviteList(w *entity.MysqlWhere) (list []*do.ChatInvite, err error)
}

type chatInviteRepository struct {
}

func NewChatInviteRepository() ChatInviteRepository {
	return &chatInviteRepository{}
}

func (r *chatInviteRepository) CreateChatInvites(invites []*po.ChatInvite) (err error) {
	db := xmysql.GetDB()
	err = db.Create(invites).Error
	return
}

func (r *chatInviteRepository) TxCreateChatInvites(tx *gorm.DB, list []*po.ChatInvite) (err error) {
	err = tx.Create(list).Error
	return
}

func (r *chatInviteRepository) TxUpdateChatInvite(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64) {
	rowsAffected = tx.Model(po.ChatInvite{}).Where(u.Query, u.Args...).Updates(u.Values).RowsAffected
	return
}

func (r *chatInviteRepository) TxChatInvite(tx *gorm.DB, w *entity.MysqlWhere) (invite *po.ChatInvite, err error) {
	invite = new(po.ChatInvite)
	err = tx.Where(w.Query, w.Args...).Find(invite).Error
	return
}

func (r *chatInviteRepository) ChatInvite(w *entity.MysqlWhere) (invite *po.ChatInvite, err error) {
	invite = new(po.ChatInvite)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(invite).Error
	return
}

func (r *chatInviteRepository) ChatInviteList(w *entity.MysqlWhere) (list []*do.ChatInvite, err error) {
	list = make([]*do.ChatInvite, 0)
	db := xmysql.GetDB()
	err = db.Model(do.ChatInvite{}).
		Preload("InitiatorInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid, lark_id, nickname, gender, birth_ts, city_id, avatar_key")
		}).
		Where(w.Query, w.Args...).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}
