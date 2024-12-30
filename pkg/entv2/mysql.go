package entity

import (
	"reflect"
	"time"
)

type GormCreatedTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime;NOT NULL" json:"created_ts"`
}

type GormUpdatedTs struct {
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime;NOT NULL" json:"updated_ts"`
}

type GormDeletedTs struct {
	DeletedTs int64 `gorm:"column:deleted_ts;default:0;NOT NULL" json:"deleted_ts"`
}

type GormEntityTs struct {
	GormCreatedTs
	GormUpdatedTs
	GormDeletedTs
}

type GormTs struct {
	GormCreatedTs
	GormUpdatedTs
}

func Deleted() (column string, value interface{}) {
	return "deleted_ts", time.Now().Unix()
}

type Mysql struct {
	Query string
	Args  []interface{}
}

func (m *Mysql) appendCondition(condition string, isAnd bool) {
	if m.Query == "" {
		m.Query = condition
	} else {
		if isAnd {
			m.Query += " AND " + condition
		} else {
			m.Query = "(" + m.Query + ") OR (" + condition + ")"
		}
	}
}

func (m *Mysql) andCondition(condition string) {
	m.appendCondition(condition, true)
}

func (m *Mysql) orCondition(condition string) {
	m.appendCondition(condition, false)
}

func (m *Mysql) NotDeleted(alias string) {
	m.andCondition(alias + ".deleted_ts=0")
}

func (m *Mysql) IsNull(field string) {
	m.andCondition(field + " IS NULL")
}

func (m *Mysql) IsNotNull(field string) {
	m.andCondition(field + " IS NOT NULL")
}

func (m *Mysql) Where(query string, value ...interface{}) {
	m.andCondition(query)
	m.Args = append(m.Args, value...)
}

func (m *Mysql) OrWhere(query string, value ...interface{}) {
	m.orCondition(query)
	m.Args = append(m.Args, value...)
}

func (m *Mysql) Between(field string, begin interface{}, end interface{}) {
	m.andCondition(field + " BETWEEN ? AND ?")
	m.Args = append(m.Args, begin, end)
}

func (m *Mysql) Bracket() {
	m.Query = "(" + m.Query + ")"
}

func (m *Mysql) OpenParen() {
	m.Query += "("
}

func (m *Mysql) CloseParen() {
	m.Query += ")"
}

func (m *Mysql) WhereQuery(query string) {
	m.andCondition(query)
}

func (m *Mysql) OrQuery(query string) {
	m.orCondition(query)
}

func (m *Mysql) AppendArg(value interface{}) {
	m.Args = append(m.Args, value)
}

func (m *Mysql) SetConditions(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag.Get("where")
		if tag != "" {
			m.Where(tag+"=?", value)
		}
	}
}

func (m *Mysql) Equal(field string, value interface{}) {
	m.Where(field+" = ?", value)
}

func (m *Mysql) NotEqual(field string, value interface{}) {
	m.Where(field+" != ?", value)
}

func (m *Mysql) Like(field string, value interface{}) {
	m.Where(field+" LIKE ?", value)
}

func (m *Mysql) NotLike(field string, value interface{}) {
	m.Where(field+" NOT LIKE ?", value)
}

func (m *Mysql) GreaterThan(field string, value interface{}) {
	m.Where(field+" > ?", value)
}

func (m *Mysql) LessThan(field string, value interface{}) {
	m.Where(field+" < ?", value)
}

func (m *Mysql) GreaterEqual(field string, value interface{}) {
	m.Where(field+" >= ?", value)
}

func (m *Mysql) LessEqual(field string, value interface{}) {
	m.Where(field+" <= ?", value)
}

func (m *Mysql) In(field string, value interface{}) {
	m.Where(field+" IN (?)", value)
}

func (m *Mysql) NotIn(field string, value interface{}) {
	m.Where(field+" NOT IN (?)", value)
}

func (m *Mysql) OrEqual(field string, value interface{}) {
	m.OrWhere(field+" = ?", value)
}

func (m *Mysql) OrNotEqual(field string, value interface{}) {
	m.OrWhere(field+" != ?", value)
}

func (m *Mysql) OrLike(field string, value interface{}) {
	m.OrWhere(field+" LIKE ?", value)
}

func (m *Mysql) OrNotLike(field string, value interface{}) {
	m.OrWhere(field+" NOT LIKE ?", value)
}

func (m *Mysql) OrGreaterThan(field string, value interface{}) {
	m.OrWhere(field+" > ?", value)
}

func (m *Mysql) OrLessThan(field string, value interface{}) {
	m.OrWhere(field+" < ?", value)
}

func (m *Mysql) OrGreaterEqual(field string, value interface{}) {
	m.OrWhere(field+" >= ?", value)
}

func (m *Mysql) OrLessEqual(field string, value interface{}) {
	m.OrWhere(field+" <= ?", value)
}

func (m *Mysql) OrIn(field string, value interface{}) {
	m.OrWhere(field+" IN (?)", value)
}

func (m *Mysql) OrNotIn(field string, value interface{}) {
	m.OrWhere(field+" NOT IN (?)", value)
}

func (m *Mysql) OrBetween(field string, begin interface{}, end interface{}) {
	m.OrWhere(field+" BETWEEN ? AND ?", begin, end)
}
