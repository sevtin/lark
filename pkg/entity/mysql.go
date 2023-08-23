package entity

import (
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
	return "deleted_ts", time.Now().UnixNano() / 1e6
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
		Query: "1=1",
		Args:  make([]interface{}, 0),
	}
}

func (m *MysqlQuery) SetFilter(query string, value ...interface{}) {
	m.Query += " AND " + query
	m.Args = append(m.Args, value...)
}

func (m *MysqlQuery) Between(field string, begin interface{}, end interface{}) {
	m.Query += " AND " + field + " BETWEEN ? AND ?"
	m.Args = append(m.Args, begin)
	m.Args = append(m.Args, end)
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
	m.Query += " AND " + query
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
	m.Query = "1=1"
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

type MysqlUpdate struct {
	Query  string
	Args   []interface{}
	Values map[string]interface{}
}

func NewMysqlUpdate() *MysqlUpdate {
	return &MysqlUpdate{
		Query:  "1=1",
		Args:   make([]interface{}, 0),
		Values: make(map[string]interface{}),
	}
}

func (m *MysqlUpdate) Set(key string, value interface{}) {
	m.Values[key] = value
}

func (m *MysqlUpdate) SetFilter(query string, value ...interface{}) {
	m.Query += " AND " + query
	m.Args = append(m.Args, value...)
}

func (m *MysqlUpdate) AndQuery(query string) {
	m.Query += " AND " + query
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
	m.Query = "1=1"
	m.Args = make([]interface{}, 0)
	m.Values = make(map[string]interface{})
}

func (m *MysqlUpdate) Updates(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Model(model).Where(m.Query, m.Args...).Updates(m.Values)
}
