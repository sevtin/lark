package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
)

type AuthRepository interface {
	Create(user *po.User) (err error)
	TxCreate(tx *gorm.DB, user *po.User) (err error)
	VerifyIdentity(w *entity.MysqlWhere) (user *po.User, err error)
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

/*
存:传指针对象，Create时不需要&，同时会Out表中的数据
读:返回指针对象，Find时不需要&
需要不为nil
*/

func (r *authRepository) Create(user *po.User) (err error) {
	user.Uid = xsnowflake.NewSnowflakeID()
	if user.LarkId == "" {
		user.LarkId = xsnowflake.DefaultLarkId()
	}
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *authRepository) TxCreate(tx *gorm.DB, user *po.User) (err error) {
	user.Uid = xsnowflake.NewSnowflakeID()
	if user.LarkId == "" {
		user.LarkId = xsnowflake.DefaultLarkId()
	}
	err = tx.Create(user).Error
	return
}

func (r *authRepository) VerifyIdentity(w *entity.MysqlWhere) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(user).Error
	return
}
