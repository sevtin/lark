package repo

import (
	"gorm.io/gorm"
	"lark/domain/pdo"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type OauthUserRepository interface {
	TxCreateOauthUser(tx *gorm.DB, user *po.OauthUser) (err error)
	TxUpdateOauthUser(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	GetOAuthUserInfo(q *entity.MysqlQuery) (user *pdo.OauthUserInfo, err error)
	GetOAuthUser(q *entity.MysqlQuery) (user *pdo.OauthUser, err error)
	UpdateOauthUser(u *entity.MysqlUpdate) (err error)
	GetUserToken(q *entity.MysqlQuery) (user *pdo.OauthUserToken, err error)
}

type oauthUserRepository struct {
}

func NewOauthUserRepository() OauthUserRepository {
	return &oauthUserRepository{}
}

func (r *oauthUserRepository) TxCreateOauthUser(tx *gorm.DB, user *po.OauthUser) (err error) {
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *oauthUserRepository) TxUpdateOauthUser(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	u.Updates(tx, &po.OauthUser{})
	return
}

func (r *oauthUserRepository) GetOAuthUser(q *entity.MysqlQuery) (user *pdo.OauthUser, err error) {
	user = new(pdo.OauthUser)
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Select(user.GetFields()).Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *oauthUserRepository) GetOAuthUserInfo(q *entity.MysqlQuery) (user *pdo.OauthUserInfo, err error) {
	user = new(pdo.OauthUserInfo)
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Select(user.GetFields()).Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *oauthUserRepository) GetUserToken(q *entity.MysqlQuery) (user *pdo.OauthUserToken, err error) {
	user = new(pdo.OauthUserToken)
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Select(user.GetFields()).Where(q.Query, q.Args...).Find(user).Error
	return
}

func (r *oauthUserRepository) UpdateOauthUser(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(&po.OauthUser{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}
