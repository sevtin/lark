package repo

import (
	"gorm.io/gorm"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_convo"
)

type ChatMemberRepository interface {
	TxCreate(tx *gorm.DB, chatMember *po.ChatMember) (err error)
	TxCreateMultiple(tx *gorm.DB, users []*po.ChatMember) (err error)
	ChatMemberStatusList(w *entity.MysqlQuery) (list []*do.ChatMemberStatus, err error)
	DistMemberList(w *entity.MysqlQuery) (list []*pb_chat_member.DistMember, err error)
	ChatMember(w *entity.MysqlQuery) (member *pb_chat_member.ChatMemberInfo, err error)
	ChatMemberCount(w *entity.MysqlQuery) (count int64, err error)
	ChatMemberList(w *entity.MysqlQuery) (members []*pb_chat_member.ChatMemberInfo, err error)
	UpdateChatMember(u *entity.MysqlUpdate) (err error)
	TxUpdateChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	TxQuitChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64, err error)
	ChatMemberBasicInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.ChatMemberBasicInfo, err error)
	GroupChatBasicInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.GroupChatBasicInfo, err error)
	GroupChatMemberInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.GroupChatMemberInfo, err error)
	ConvoChatSeqList(q *entity.MysqlQuery) (list []*pb_convo.ConvoChatSeq, err error)
}

type chatMemberRepository struct {
}

func NewChatMemberRepository() ChatMemberRepository {
	return &chatMemberRepository{}
}

func (r *chatMemberRepository) TxCreate(tx *gorm.DB, chatMember *po.ChatMember) (err error) {
	err = tx.Create(chatMember).Error
	return
}

func (r *chatMemberRepository) TxCreateMultiple(tx *gorm.DB, users []*po.ChatMember) (err error) {
	err = tx.Create(users).Error
	return
}

func (r *chatMemberRepository) ChatMemberStatusList(w *entity.MysqlQuery) (list []*do.ChatMemberStatus, err error) {
	list = make([]*do.ChatMemberStatus, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).
		Select("chat_id,status").
		Where(w.Query, w.Args...).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}

func (r *chatMemberRepository) DistMemberList(w *entity.MysqlQuery) (list []*pb_chat_member.DistMember, err error) {
	list = make([]*pb_chat_member.DistMember, 0)
	db := xmysql.GetDB()
	err = db.Table("chat_members m").
		Select("u.server_id,m.uid,m.status").
		Joins("LEFT JOIN users u ON u.uid=m.uid").
		Where(w.Query, w.Args...).
		Find(&list).Error
	return
}

func (r *chatMemberRepository) ChatMember(w *entity.MysqlQuery) (member *pb_chat_member.ChatMemberInfo, err error) {
	member = new(pb_chat_member.ChatMemberInfo)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("chat_id,chat_type,uid,alias,member_avatar,role_id,status").Where(w.Query, w.Args...).Find(&member).Error
	return
}

func (r *chatMemberRepository) ChatMemberCount(w *entity.MysqlQuery) (count int64, err error) {
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(w.Query, w.Args...).Count(&count).Error
	return
}

func (r *chatMemberRepository) ChatMemberList(w *entity.MysqlQuery) (members []*pb_chat_member.ChatMemberInfo, err error) {
	members = make([]*pb_chat_member.ChatMemberInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).
		Select("chat_id,chat_type,uid,alias,member_avatar,role_id").
		Where(w.Query, w.Args...).
		Limit(w.Limit).
		Find(&members).Error
	return
}

func (r *chatMemberRepository) UpdateChatMember(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatMemberRepository) TxUpdateChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatMemberRepository) TxQuitChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (rowsAffected int64, err error) {
	result := tx.Model(po.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values)
	err = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (r *chatMemberRepository) ChatMemberBasicInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.ChatMemberBasicInfo, err error) {
	list = make([]*pb_chat_member.ChatMemberBasicInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("uid,alias,remark,member_avatar,status").
		Where(w.Query, w.Args...).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}

func (r *chatMemberRepository) GroupChatMemberInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.GroupChatMemberInfo, err error) {
	list = make([]*pb_chat_member.GroupChatMemberInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("uid,alias,member_avatar,role_id,status").
		Where(w.Query, w.Args...).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}

func (r *chatMemberRepository) GroupChatBasicInfoList(w *entity.MysqlQuery) (list []*pb_chat_member.GroupChatBasicInfo, err error) {
	list = make([]*pb_chat_member.GroupChatBasicInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("chat_id,chat_name,remark,chat_avatar").
		Where(w.Query, w.Args...).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}

func (r *chatMemberRepository) ConvoChatSeqList(q *entity.MysqlQuery) (list []*pb_convo.ConvoChatSeq, err error) {
	list = make([]*pb_convo.ConvoChatSeq, 0)
	db := xmysql.GetDB()
	err = db.Table("chat_members m").
		Select("m.read_seq,c.chat_id,c.seq_id,c.srv_ts").
		Joins("LEFT JOIN chats c ON c.chat_id=m.chat_id").
		Where(q.Query, q.Args...).
		Order("c.srv_ts DESC").
		Limit(q.Limit).
		Find(&list).Error
	return
}
