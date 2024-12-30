package entity

import (
	"gorm.io/gorm"
	"reflect"
)

type MysqlQuery struct {
	Mysql
	Model  interface{}
	Fields string
	Limit  int
	Offset int
	Sort   string
}

func NewMysqlQuery() *MysqlQuery {
	return &MysqlQuery{
		Mysql: Mysql{
			Query: "deleted_ts=0",
			Args:  make([]interface{}, 0),
		},
	}
}

func NewNormalQuery() *MysqlQuery {
	return &MysqlQuery{
		Mysql: Mysql{
			Args: make([]interface{}, 0),
		},
	}
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
	m.Reset()
	m.Query = ""
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
