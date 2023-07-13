package mrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
)

type MessageHotRepository interface {
	Create(message *po.Message) (err error)
	Update(u *entity.MongoUpdate) (err error)
	Messages(w *entity.MongoQuery) (messages []*po.Message, err error)
}

type messageHotRepository struct {
}

func NewMessageHotRepository() MessageHotRepository {
	return &messageHotRepository{}
}

func (r *messageHotRepository) Create(message *po.Message) (err error) {
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_MESSAGES)
	defer cancel()
	if coll == nil {
		return
	}
	if _, err = coll.InsertOne(ctx, message); err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}

func (r *messageHotRepository) Messages(w *entity.MongoQuery) (messages []*po.Message, err error) {
	messages = make([]*po.Message, 0)
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
		cur    *mongo.Cursor
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_MESSAGES)
	defer cancel()
	if coll == nil {
		return
	}
	cur, err = coll.Find(ctx, w.Filter, w.FindOptions)
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &messages)
	return
}

func (r *messageHotRepository) Update(u *entity.MongoUpdate) (err error) {
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_MESSAGES)
	defer cancel()
	if coll == nil {
		return
	}
	if _, err = coll.UpdateMany(ctx, u.Filter, u.Update); err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}
