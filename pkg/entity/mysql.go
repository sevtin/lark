package entity

import "time"

type GormCreatedTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli;NOT NULL" json:"created_ts"`
}

type GormUpdatedTs struct {
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli;NOT NULL" json:"updated_ts"`
}

type GormDeletedTs struct {
	DeletedTs int64 `gorm:"column:deleted_ts;default:0;NOT NULL" json:"deleted_ts"`
}

type GormEntityTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli;NOT NULL" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli;NOT NULL" json:"updated_ts"`
	DeletedTs int64 `gorm:"column:deleted_ts;default:0;NOT NULL" json:"deleted_ts"`
}

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli;NOT NULL" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli;NOT NULL" json:"updated_ts"`
}

func Deleted() (column string, value interface{}) {
	return "deleted_ts", time.Now().UnixNano() / 1e6
}

type MysqlWhere struct {
	Query  string
	Args   []interface{}
	Limit  int
	Offset int
	Sort   string
}

func NewMysqlWhere() *MysqlWhere {
	return &MysqlWhere{
		Query: "deleted_ts=0",
		Args:  make([]interface{}, 0),
	}
}

func NewNormalWhere() *MysqlWhere {
	return &MysqlWhere{
		Query: "1=1",
		Args:  make([]interface{}, 0),
	}
}

func (m *MysqlWhere) SetFilter(query string, value interface{}) {
	m.Query += " AND " + query
	m.Args = append(m.Args, value)
}

func (m *MysqlWhere) SetSort(sort string) {
	m.Sort = sort
}

func (m *MysqlWhere) SetOffset(offset int) {
	m.Offset = offset
}

func (m *MysqlWhere) SetLimit(limit int32) {
	m.Limit = int(limit)
}

func (m *MysqlWhere) AndQuery(query string) {
	m.Query += " AND " + query
}

func (m *MysqlWhere) AppendArg(value interface{}) {
	m.Args = append(m.Args, value)
}

func (m *MysqlWhere) Reset() {
	m.Query = "deleted_ts=0"
	m.Args = make([]interface{}, 0)
	m.Sort = ""
	m.Limit = 0
	m.Offset = 0
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

func (m *MysqlUpdate) SetFilter(query string, value interface{}) {
	m.Query += " AND " + query
	m.Args = append(m.Args, value)
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
