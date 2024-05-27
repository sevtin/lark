package entity

import (
	"fmt"
	"gorm.io/gorm"
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
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime;NOT NULL" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime;NOT NULL" json:"updated_ts"`
	DeletedTs int64 `gorm:"column:deleted_ts;default:0;NOT NULL" json:"deleted_ts"`
}

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime;NOT NULL" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime;NOT NULL" json:"updated_ts"`
}

func Deleted() (column string, value interface{}) {
	return "deleted_ts", time.Now().Unix()
}

type MysqlQuery struct {
	Model  interface{}
	Fields string
	Query  string
	Args   []interface{}
	Limit  int
	Offset int
	Sort   string
}

func NewMysqlQuery() *MysqlQuery {
	return &MysqlQuery{
		Query: "deleted_ts=0",
		Args:  make([]interface{}, 0),
	}
}

func NewNormalQuery() *MysqlQuery {
	return &MysqlQuery{
		Args: make([]interface{}, 0),
	}
}

func (m *MysqlQuery) andCondition(condition string) {
	if m.Query == "" {
		m.Query = condition
	} else {
		m.Query += " AND " + condition
	}
}

func (m *MysqlQuery) orCondition(condition string) {
	if m.Query == "" {
		m.Query = condition
	} else {
		m.Query += " Or " + condition
	}
}

func (m *MysqlQuery) NotDeleted(alias string) {
	clause := alias + ".deleted_ts=0"
	m.andCondition(clause)
}

func (m *MysqlQuery) IsNull(field string) {
	clause := field + " IS NULL"
	m.andCondition(clause)
}

func (m *MysqlQuery) IsNotNull(field string) {
	clause := field + " IS NOT NULL"
	m.andCondition(clause)
}

func (m *MysqlQuery) SetFilter(query string, value ...interface{}) {
	m.andCondition(query)
	m.Args = append(m.Args, value...)
}

func (m *MysqlQuery) SetFilterOr(query string, value ...interface{}) {
	m.orCondition(query)
	m.Args = append(m.Args, value...)
}

func (m *MysqlQuery) Between(field string, begin interface{}, end interface{}) {
	clause := field + " BETWEEN ? AND ?"
	m.andCondition(clause)
	m.Args = append(m.Args, begin)
	m.Args = append(m.Args, end)
}

func (m *MysqlQuery) OpenParen() {
	m.Query += "(1=1"
}

func (m *MysqlQuery) CloseParen() {
	m.Query += ")"
}

func (m *MysqlQuery) SetSort(sort string) {
	m.Sort = sort
}

func (m *MysqlQuery) SetOffset(offset int) {
	m.Offset = offset
}

func (m *MysqlQuery) SetLimit(limit int32) {
	m.Limit = int(limit)
}

func (m *MysqlQuery) AndQuery(query string) {
	m.andCondition(query)
}

func (m *MysqlQuery) OrQuery(query string) {
	m.orCondition(query)
}

func (m *MysqlQuery) SetLink(link string) {
	m.Query += fmt.Sprintf(" %s ", link)
}

func (m *MysqlQuery) AppendArg(value interface{}) {
	m.Args = append(m.Args, value)
}

func (m *MysqlQuery) Reset() {
	m.Model = nil
	m.Fields = ""
	m.Query = "deleted_ts=0"
	m.Args = make([]interface{}, 0)
	m.Sort = ""
	m.Limit = 0
	m.Offset = 0
}

func (m *MysqlQuery) Normal() {
	m.Model = nil
	m.Fields = ""
	m.Query = ""
	m.Args = make([]interface{}, 0)
	m.Sort = ""
	m.Limit = 0
	m.Offset = 0
}

func (m *MysqlQuery) Find(db *gorm.DB, dist interface{}) (err error) {
	err = db.Model(m.Model).Select(m.Fields).Where(m.Query, m.Args...).Find(dist).Error
	return
}

func (m *MysqlQuery) Finds(db *gorm.DB, model any) (dest any, err error) {
	var (
		typ        = reflect.TypeOf(model)
		sliceType  = reflect.SliceOf(typ)
		sliceValue = reflect.New(sliceType).Elem()
	)
	dest = sliceValue.Interface()
	err = db.Model(m.Model).Select(m.Fields).Where(m.Query, m.Args...).Find(dest).Error
	return
}

func (m *MysqlQuery) Delete(db *gorm.DB) (err error) {
	return db.Model(m.Model).Where(m.Query, m.Args...).Update(Deleted()).Error
}

func (m *MysqlQuery) SetConditions(obj any) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag.Get("where")
		if tag != "" {
			m.SetFilter(tag+"=?", value)
		}
	}
}

type MysqlUpdate struct {
	Query  string
	Args   []interface{}
	Values map[string]interface{}
}

func NewMysqlUpdate() *MysqlUpdate {
	return &MysqlUpdate{
		Args:   make([]interface{}, 0),
		Values: make(map[string]interface{}),
	}
}

func (m *MysqlUpdate) andCondition(condition string) {
	if m.Query == "" {
		m.Query = condition
	} else {
		m.Query += " AND " + condition
	}
}

func (m *MysqlUpdate) orCondition(condition string) {
	if m.Query == "" {
		m.Query = condition
	} else {
		m.Query += " Or " + condition
	}
}

func (m *MysqlUpdate) Set(key string, value interface{}) {
	m.Values[key] = value
}

func (m *MysqlUpdate) SetFilter(query string, value ...interface{}) {
	m.andCondition(query)
	m.Args = append(m.Args, value...)
}

func (m *MysqlUpdate) AndQuery(query string) {
	m.andCondition(query)
}

func (m *MysqlUpdate) AppendArg(value interface{}) {
	m.Args = append(m.Args, value)
}

func (m *MysqlUpdate) Reset() {
	m.Query = "deleted_ts=0"
	m.Args = make([]interface{}, 0)
	m.Values = make(map[string]interface{})
}

func (m *MysqlUpdate) Normal() {
	m.Query = ""
	m.Args = make([]interface{}, 0)
	m.Values = make(map[string]interface{})
}

func (m *MysqlUpdate) Updates(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Model(model).Where(m.Query, m.Args...).Updates(m.Values)
}

func (m *MysqlUpdate) Delete(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Model(model).Where(m.Query, m.Args...).Update(Deleted())
}

func (m *MysqlUpdate) SetConditionsAndValues(obj any) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag.Get("where")
		if tag != "" {
			m.SetFilter(tag+"=?", value)
		}
		tag = field.Tag.Get("update")
		if tag != "" {
			m.Set(tag, value)
		}
	}
}
