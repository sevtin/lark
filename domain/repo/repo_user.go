package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

type UserRepository interface {
	Create(user *po.User) (err error)
	VerifyUserIdentity(w *entity.MysqlWhere) (user *po.User, err error)
	Exists(w *entity.MysqlWhere, uid int64) (exists bool, err error)
	TxExists(tx *gorm.DB, w *entity.MysqlWhere, uid int64) (exists bool, err error)
	UserInfo(w *entity.MysqlWhere) (user *po.User, err error)
	BasicUserInfo(w *entity.MysqlWhere) (user *pb_user.BasicUserInfo, err error)
	BasicUserInfoList(w *entity.MysqlWhere) (list []*pb_user.BasicUserInfo, err error)
	UserList(w *entity.MysqlWhere) (list []*po.User, err error)
	TxUserList(tx *gorm.DB, w *entity.MysqlWhere) (list []*po.User, err error)
	TxBasicUserList(tx *gorm.DB, w *entity.MysqlWhere) (list []*pb_user.BasicUserInfo, err error)
	TxUserSrvList(tx *gorm.DB, w *entity.MysqlWhere) (list []*pb_user.UserSrvInfo, err error)
	UpdateUser(u *entity.MysqlUpdate) (err error)
	TxUpdateUser(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	UserServerList(u *entity.MysqlWhere) (list []*pb_user.UserServerId, err error)
	UserServerId(u *entity.MysqlWhere) (server *pb_user.UserServerId, err error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

/*
存:传指针对象，Create时不需要&，同时会Out表中的数据
读:返回指针对象，Find时不需要&，需要不为nil
*/

func (r *userRepository) Create(user *po.User) (err error) {
	user.Uid = xsnowflake.NewSnowflakeID()
	if user.LarkId == "" {
		user.LarkId = xsnowflake.DefaultLarkId()
	}
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *userRepository) VerifyUserIdentity(w *entity.MysqlWhere) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(user).Error
	return
}

func (r *userRepository) UserList(w *entity.MysqlWhere) (list []*po.User, err error) {
	list = make([]*po.User, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) TxUserList(tx *gorm.DB, w *entity.MysqlWhere) (list []*po.User, err error) {
	list = make([]*po.User, 0)
	err = tx.Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) TxBasicUserList(tx *gorm.DB, w *entity.MysqlWhere) (list []*pb_user.BasicUserInfo, err error) {
	list = make([]*pb_user.BasicUserInfo, 0)
	err = tx.Model(po.User{}).Select("uid,lark_id,nickname,gender,birth_ts,city_id,avatar_key").Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) TxUserSrvList(tx *gorm.DB, w *entity.MysqlWhere) (list []*pb_user.UserSrvInfo, err error) {
	list = make([]*pb_user.UserSrvInfo, 0)
	err = tx.Model(po.User{}).Select("uid,nickname,avatar_key,server_id").Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) Exists(w *entity.MysqlWhere, uid int64) (exists bool, err error) {
	var (
		user = new(po.User)
		db   = xmysql.GetDB()
	)
	err = db.Select("uid").Where(w.Query, w.Args...).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if user.Uid > 0 && user.Uid != uid {
		exists = true
	}
	return
}

func (r *userRepository) TxExists(tx *gorm.DB, w *entity.MysqlWhere, uid int64) (exists bool, err error) {
	var (
		user = new(po.User)
	)
	err = tx.Select("uid").Where(w.Query, w.Args...).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if user.Uid > 0 && user.Uid != uid {
		exists = true
	}
	return
}

func (r *userRepository) UserInfo(w *entity.MysqlWhere) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&user).Error
	return
}

func (r *userRepository) BasicUserInfo(w *entity.MysqlWhere) (user *pb_user.BasicUserInfo, err error) {
	user = new(pb_user.BasicUserInfo)
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Select("uid,lark_id,nickname,gender,birth_ts,city_id,avatar_key").Where(w.Query, w.Args...).Find(&user).Error
	return
}

func (r *userRepository) BasicUserInfoList(w *entity.MysqlWhere) (list []*pb_user.BasicUserInfo, err error) {
	list = make([]*pb_user.BasicUserInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Select("uid,lark_id,nickname,gender,birth_ts,city_id,avatar_key").Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) UpdateUser(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *userRepository) TxUpdateUser(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.User{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *userRepository) UserServerList(u *entity.MysqlWhere) (list []*pb_user.UserServerId, err error) {
	list = make([]*pb_user.UserServerId, 0)
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Select("uid,server_id").Where(u.Query, u.Args...).Find(&list).Error
	return
}

func (r *userRepository) UserServerId(u *entity.MysqlWhere) (server *pb_user.UserServerId, err error) {
	server = new(pb_user.UserServerId)
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Select("uid,server_id").Where(u.Query, u.Args...).Find(server).Error
	return
}
