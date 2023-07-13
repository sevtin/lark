package entity

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lark/pkg/common/xmongo"
	"time"
)

func Collection(collection string) (ctx context.Context, cancel context.CancelFunc, coll *mongo.Collection) {
	var (
		db *mongo.Database
	)
	ctx, cancel = NewContext()
	db = xmongo.GetDB()
	if db == nil {
		return
	}
	coll = db.Collection(collection)
	return
}

func NewContext() (ctx context.Context, cancelFunc context.CancelFunc) {
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	return
}

type MongoQuery struct {
	Filter      bson.D
	FindOptions *options.FindOptions
}

func NewMongoQuery() *MongoQuery {
	return &MongoQuery{
		Filter:      bson.D{},
		FindOptions: new(options.FindOptions),
	}
}

//func (m *MongoQuery) SetSort(key string, asc bool) {
//	var val = 1
//	if asc == false {
//		val = -1
//	}
//	m.FindOptions.SetSort(bson.D{bson.E{key, val}})
//}

func (m *MongoQuery) SetSort(sort bson.D) {
	m.FindOptions.SetSort(sort)
}

func (m *MongoQuery) SetLimit(limit int64) {
	m.FindOptions.SetLimit(limit)
}

func (m *MongoQuery) SetSkip(skip int64) {
	m.FindOptions.SetSkip(skip)
}

func (m *MongoQuery) SetFilter(key string, value interface{}) {
	m.Filter = append(m.Filter, bson.D{{key, value}}...)
}

type MongoUpdate struct {
	Filter bson.D
	Update bson.D
}

func NewMongoUpdate() *MongoUpdate {
	return &MongoUpdate{
		Filter: bson.D{},
		Update: bson.D{},
	}
}

func (m *MongoUpdate) SetFilter(key string, value interface{}) {
	m.Filter = append(m.Filter, bson.D{{key, value}}...)
}

func (m *MongoUpdate) Set(key string, value interface{}) {
	m.Update = append(m.Update, bson.D{{"$set", bson.D{{key, value}}}}...)
}
