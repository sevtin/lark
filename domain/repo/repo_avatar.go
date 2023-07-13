package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type AvatarRepository interface {
	Avatar(w *entity.MysqlQuery) (avatar *po.Avatar, err error)
	AvatarList(w *entity.MysqlQuery) (avatars []*po.Avatar, err error)
	TxCreate(tx *gorm.DB, avatar *po.Avatar) (err error)
	TxUpdateAvatar(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
}

type avatarRepository struct {
}

func NewAvatarRepository() AvatarRepository {
	return &avatarRepository{}
}

func (r *avatarRepository) Avatar(w *entity.MysqlQuery) (avatar *po.Avatar, err error) {
	avatar = new(po.Avatar)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(avatar).Error
	return
}

func (r *avatarRepository) AvatarList(w *entity.MysqlQuery) (avatars []*po.Avatar, err error) {
	avatars = make([]*po.Avatar, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&avatars).Error
	return
}

func (r *avatarRepository) TxUpdateAvatar(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.Avatar{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *avatarRepository) TxCreate(tx *gorm.DB, avatar *po.Avatar) (err error) {
	err = tx.Create(avatar).Error
	return
}
