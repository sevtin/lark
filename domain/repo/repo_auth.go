package repo

import (
	"gorm.io/gorm"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
)

type AuthRepository interface {
	Create(user *po.User) (err error)
	TxCreate(tx *gorm.DB, user *po.User) (err error)
	VerifyIdentity(q *entity.MysqlQuery) (user *po.User, err error)
	TxCreateOauthUser(tx *gorm.DB, user *po.OauthUser) (err error)
	GetOAuthUser(q *entity.MysqlQuery) (user *pdo.OauthUser, err error)
	UpdateOauthUser(u *entity.MysqlUpdate) (err error)
	GetUserToken(q *entity.MysqlQuery) (user *pdo.OauthUserToken, err error)
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
	err = tx.Create(user).Error
	return
}

func (r *authRepository) VerifyIdentity(q *entity.MysqlQuery) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *authRepository) TxCreateOauthUser(tx *gorm.DB, user *po.OauthUser) (err error) {
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *authRepository) GetOAuthUser(q *entity.MysqlQuery) (user *pdo.OauthUser, err error) {
	user = new(pdo.OauthUser)
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Select(user.GetFields()).Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *authRepository) GetUserToken(q *entity.MysqlQuery) (user *pdo.OauthUserToken, err error) {
	user = new(pdo.OauthUserToken)
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Select(user.GetFields()).Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *authRepository) UpdateOauthUser(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}
