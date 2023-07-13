package mrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
)

type LbsRepository interface {
	Upsert(location *po.UserLocation) (err error)
	Update(u *entity.MongoUpdate) (err error)
	UserLocations(w *entity.MongoQuery) (locations []*po.UserLocation, err error)
}

type lbsRepository struct {
}

func NewLbsRepository() LbsRepository {
	return &lbsRepository{}
}

func (r *lbsRepository) Upsert(location *po.UserLocation) (err error) {
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
		filter = bson.D{{"uid", location.Uid}}
		update = bson.D{{"$set", location}}
		opt    = options.Update().SetUpsert(true)
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_USER_LOCATIONS)
	defer cancel()
	if coll == nil {
		return
	}
	if _, err = coll.UpdateOne(ctx, filter, update, opt); err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}

func (r *lbsRepository) UserLocations(w *entity.MongoQuery) (locations []*po.UserLocation, err error) {
	locations = make([]*po.UserLocation, 0)
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
		cur    *mongo.Cursor
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_USER_LOCATIONS)
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
	err = cur.All(ctx, &locations)
	return
}

func (r *lbsRepository) Update(u *entity.MongoUpdate) (err error) {
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel, coll = entity.Collection(po.MONGO_COLLECTION_USER_LOCATIONS)
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
