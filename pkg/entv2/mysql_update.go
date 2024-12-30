package entity

import (
	"gorm.io/gorm"
	"reflect"
)

type MysqlUpdate struct {
	Mysql
	Values map[string]interface{}
}

func NewMysqlUpdate() *MysqlUpdate {
	return &MysqlUpdate{
		Mysql: Mysql{
			Args: make([]interface{}, 0),
		},
		Values: make(map[string]interface{}),
	}
}

func (m *MysqlUpdate) Set(key string, value interface{}) {
	m.Values[key] = value
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
			m.Where(tag+"=?", value)
		}
		tag = field.Tag.Get("update")
		if tag != "" {
			m.Set(tag, value)
		}
	}
}
