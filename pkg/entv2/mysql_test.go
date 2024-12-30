package entity

import (
	"reflect"
	"testing"
)

func TestMysql(t *testing.T) {
	tests := []struct {
		name      string
		testFunc  func(m *Mysql)
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name: "Test NotDeleted",
			testFunc: func(m *Mysql) {
				m.NotDeleted("orders")
			},
			wantQuery: "orders.deleted_ts=0",
			wantArgs:  nil,
		},
		{
			name: "Test IsNull",
			testFunc: func(m *Mysql) {
				m.IsNull("orders.order_status")
			},
			wantQuery: "orders.order_status IS NULL",
			wantArgs:  nil,
		},
		{
			name: "Test IsNotNull",
			testFunc: func(m *Mysql) {
				m.IsNotNull("orders.order_status")
			},
			wantQuery: "orders.order_status IS NOT NULL",
			wantArgs:  nil,
		},
		{
			name: "Test Where",
			testFunc: func(m *Mysql) {
				m.Where("orders.order_id = ?", 12345)
			},
			wantQuery: "orders.order_id = ?",
			wantArgs:  []interface{}{12345},
		},
		{
			name: "Test OrWhere",
			testFunc: func(m *Mysql) {
				m.OrWhere("orders.uid = ?", 67890)
			},
			wantQuery: "orders.uid = ?",
			wantArgs:  []interface{}{67890},
		},
		{
			name: "Test Between",
			testFunc: func(m *Mysql) {
				m.Between("orders.price", 100, 200)
			},
			wantQuery: "orders.price BETWEEN ? AND ?",
			wantArgs:  []interface{}{100, 200},
		},
		{
			name: "Test Equal",
			testFunc: func(m *Mysql) {
				m.Equal("orders.order_status", 1)
			},
			wantQuery: "orders.order_status = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "Test NotEqual",
			testFunc: func(m *Mysql) {
				m.NotEqual("orders.order_status", 0)
			},
			wantQuery: "orders.order_status != ?",
			wantArgs:  []interface{}{0},
		},
		{
			name: "Test Like",
			testFunc: func(m *Mysql) {
				m.Like("orders.uid", "%123%")
			},
			wantQuery: "orders.uid LIKE ?",
			wantArgs:  []interface{}{"%123%"},
		},
		{
			name: "Test NotLike",
			testFunc: func(m *Mysql) {
				m.NotLike("orders.uid", "%456%")
			},
			wantQuery: "orders.uid NOT LIKE ?",
			wantArgs:  []interface{}{"%456%"},
		},
		{
			name: "Test GreaterThan",
			testFunc: func(m *Mysql) {
				m.GreaterThan("orders.price", 100)
			},
			wantQuery: "orders.price > ?",
			wantArgs:  []interface{}{100},
		},
		{
			name: "Test LessThan",
			testFunc: func(m *Mysql) {
				m.LessThan("orders.price", 200)
			},
			wantQuery: "orders.price < ?",
			wantArgs:  []interface{}{200},
		},
		{
			name: "Test GreaterEqual",
			testFunc: func(m *Mysql) {
				m.GreaterEqual("orders.price", 150)
			},
			wantQuery: "orders.price >= ?",
			wantArgs:  []interface{}{150},
		},
		{
			name: "Test LessEqual",
			testFunc: func(m *Mysql) {
				m.LessEqual("orders.price", 250)
			},
			wantQuery: "orders.price <= ?",
			wantArgs:  []interface{}{250},
		},
		{
			name: "Test In",
			testFunc: func(m *Mysql) {
				m.In("orders.order_id", []int{1, 2, 3})
			},
			wantQuery: "orders.order_id IN (?)",
			wantArgs:  []interface{}{[]int{1, 2, 3}},
		},
		{
			name: "Test NotIn",
			testFunc: func(m *Mysql) {
				m.NotIn("orders.order_id", []int{4, 5, 6})
			},
			wantQuery: "orders.order_id NOT IN (?)",
			wantArgs:  []interface{}{[]int{4, 5, 6}},
		},
		{
			name: "Test OrEqual",
			testFunc: func(m *Mysql) {
				m.OrEqual("orders.uid", 123)
			},
			wantQuery: "orders.uid = ?",
			wantArgs:  []interface{}{123},
		},
		{
			name: "Test OrNotEqual",
			testFunc: func(m *Mysql) {
				m.OrNotEqual("orders.uid", 456)
			},
			wantQuery: "orders.uid != ?",
			wantArgs:  []interface{}{456},
		},
		{
			name: "Test OrLike",
			testFunc: func(m *Mysql) {
				m.OrLike("orders.uid", "%789%")
			},
			wantQuery: "orders.uid LIKE ?",
			wantArgs:  []interface{}{"%789%"},
		},
		{
			name: "Test OrNotLike",
			testFunc: func(m *Mysql) {
				m.OrNotLike("orders.uid", "%000%")
			},
			wantQuery: "orders.uid NOT LIKE ?",
			wantArgs:  []interface{}{"%000%"},
		},
		{
			name: "Test OrGreaterThan",
			testFunc: func(m *Mysql) {
				m.OrGreaterThan("orders.price", 300)
			},
			wantQuery: "orders.price > ?",
			wantArgs:  []interface{}{300},
		},
		{
			name: "Test OrLessThan",
			testFunc: func(m *Mysql) {
				m.OrLessThan("orders.price", 400)
			},
			wantQuery: "orders.price < ?",
			wantArgs:  []interface{}{400},
		},
		{
			name: "Test OrGreaterEqual",
			testFunc: func(m *Mysql) {
				m.OrGreaterEqual("orders.price", 500)
			},
			wantQuery: "orders.price >= ?",
			wantArgs:  []interface{}{500},
		},
		{
			name: "Test OrLessEqual",
			testFunc: func(m *Mysql) {
				m.OrLessEqual("orders.price", 600)
			},
			wantQuery: "orders.price <= ?",
			wantArgs:  []interface{}{600},
		},
		{
			name: "Test OrIn",
			testFunc: func(m *Mysql) {
				m.OrIn("orders.market_id", []int{7, 8, 9})
			},
			wantQuery: "orders.market_id IN (?)",
			wantArgs:  []interface{}{[]int{7, 8, 9}},
		},
		{
			name: "Test OrNotIn",
			testFunc: func(m *Mysql) {
				m.OrNotIn("orders.market_id", []int{10, 11, 12})
			},
			wantQuery: "orders.market_id NOT IN (?)",
			wantArgs:  []interface{}{[]int{10, 11, 12}},
		},
		{
			name: "Test OrBetween",
			testFunc: func(m *Mysql) {
				m.OrBetween("orders.fee_rate", 0, 10)
			},
			wantQuery: "orders.fee_rate BETWEEN ? AND ?",
			wantArgs:  []interface{}{0, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mysql{}
			tt.testFunc(m)
			if m.Query != tt.wantQuery {
				t.Errorf("got query %v, want %v", m.Query, tt.wantQuery)
			}
			if !reflect.DeepEqual(m.Args, tt.wantArgs) {
				t.Errorf("got args %v, want %v", m.Args, tt.wantArgs)
			}
		})
	}
}
