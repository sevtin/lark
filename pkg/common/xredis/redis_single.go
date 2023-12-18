package xredis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"lark/pkg/conf"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"log/slog"
	"time"
)

type RedisSingle struct {
	client   *redis.Client
	redsSync *redsync.Redsync
	prefix   string
	single   bool
}

func NewRedisSingle(cfg *conf.Redis) RedisIface {
	var (
		client   *redis.Client
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// 单机redis
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Address[0],
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	// 判断是否能够链接到redis
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	// redis 锁
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)

	cs := &RedisSingle{client, redsSync, cfg.Prefix, true}
	return cs
}

func (r *RedisSingle) Single() bool {
	return r.single
}

func (r *RedisSingle) RealKey(key string) string {
	return r.prefix + key
}

func (r *RedisSingle) GetPrefix() string {
	return r.prefix
}

func (r *RedisSingle) GetClient() *redis.ClusterClient {
	return nil
}

func (r *RedisSingle) GetSingleClient() *redis.Client {
	return r.client
}

func (r *RedisSingle) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}

func (r *RedisSingle) Unlink(key string) error {
	key = RealKey(key)
	return r.client.Unlink(context.Background(), key).Err()
}

func (r *RedisSingle) TTL(key string) time.Duration {
	key = RealKey(key)
	return r.client.TTL(context.Background(), key).Val()
}

func (r *RedisSingle) Del(key string) error {
	key = RealKey(key)
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisSingle) CUnlink(keys []string) (err error) {
	var (
		pipe = r.client.Pipeline()
		key  string
	)
	for _, key = range keys {
		pipe.Unlink(context.Background(), RealKey(key))
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (r *RedisSingle) KeyExists(key string) (ok bool) {
	key = RealKey(key)
	val := r.client.Exists(context.Background(), key).Val()
	if val == 1 {
		ok = true
	}
	return
}

func (r *RedisSingle) Set(key string, value interface{}, expire time.Duration) error {
	key = RealKey(key)
	if expire > 0 {
		return r.client.Set(context.Background(), key, value, expire).Err()
	}
	return r.client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisSingle) CSet(keys []string, values []interface{}, expire time.Duration) (err error) {
	if len(keys) != len(values) {
		return
	}
	var (
		i    int
		key  string
		pipe = r.client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.Set(context.Background(), key, values[i], expire)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (r *RedisSingle) CSets(keys []string, values []interface{}, expires []time.Duration) (err error) {
	if len(keys) != len(values) || len(values) != len(expires) {
		return
	}
	var (
		i    int
		key  string
		pipe = r.client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.Set(context.Background(), key, values[i], expires[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (r *RedisSingle) Expire(key string, expire time.Duration) error {
	key = RealKey(key)
	return r.client.Expire(context.Background(), key, expire).Err()
}

func (r *RedisSingle) Get(key string) (val string, err error) {
	key = RealKey(key)
	val, err = r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisSingle) MGet(keys []string) ([]interface{}, error) {
	return r.client.MGet(context.Background(), keys...).Result()
}

func (r *RedisSingle) CMGet(keys []string) (list []string, err error) {
	var (
		key     string
		cmdList = make([]*redis.StringCmd, len(keys))
		i       int
		c       *redis.StringCmd
		pipe    = r.client.Pipeline()
	)
	for i, key = range keys {
		cmdList[i] = pipe.Get(context.Background(), RealKey(key))
	}
	_, err = pipe.Exec(context.Background())
	if err != nil {
		return
	}
	list = make([]string, len(keys))
	for i, c = range cmdList {
		list[i] = c.Val()
	}
	return
}

func (r *RedisSingle) SlotMGet(maps map[uint16][]string) (list []interface{}, err error) {
	var (
		index   = -1
		keys    []string
		cmdList = make([]*redis.SliceCmd, len(maps))
		c       *redis.SliceCmd
		pipe    = r.client.Pipeline()
	)
	for _, keys = range maps {
		c = pipe.MGet(context.Background(), keys...)
		index++
		cmdList[index] = c
	}
	_, err = pipe.Exec(context.Background())
	if err != nil {
		return
	}
	list = make([]interface{}, 0)
	for _, c = range cmdList {
		list = append(list, c.Val()...)
	}
	return
}

func (r *RedisSingle) MSet(values ...interface{}) error {
	// MSET 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。
	return r.client.MSet(context.Background(), values...).Err()
}

func (r *RedisSingle) Incr(key string) (int64, error) {
	key = RealKey(key)
	return r.client.Incr(context.Background(), key).Result()
}

func (r *RedisSingle) Decr(key string) (int64, error) {
	key = RealKey(key)
	return r.client.Decr(context.Background(), key).Result()
}

func (r *RedisSingle) GetUint64(key string) (val uint64, err error) {
	key = RealKey(key)
	val, err = r.client.Get(context.Background(), key).Uint64()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisSingle) GetInt(key string) (val int, err error) {
	key = RealKey(key)
	val, err = r.client.Get(context.Background(), key).Int()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisSingle) HGetInt64(key, field string) (value int64, err error) {
	key = RealKey(key)
	return r.client.HGet(context.Background(), key, field).Int64()
}

func (r *RedisSingle) HGetAll(key string) map[string]string {
	key = RealKey(key)
	hash := r.client.HGetAll(context.Background(), key).Val()
	return hash
}

func (r *RedisSingle) HLen(key string) int64 {
	key = RealKey(key)
	return r.client.HLen(context.Background(), key).Val()
}

func (r *RedisSingle) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	key = RealKey(key)
	return r.client.HScan(context.Background(), key, cursor, match, count).Result()
}

func (r *RedisSingle) HSet(key string, value interface{}) error {
	key = RealKey(key)
	return r.client.HSet(context.Background(), key, value).Err()
}

func (r *RedisSingle) HSetNX(key, field string, value interface{}) error {
	key = RealKey(key)
	return r.client.HSetNX(context.Background(), key, field, value).Err()
}

func (r *RedisSingle) HDels(key string, fields []string) error {
	key = RealKey(key)
	return r.client.HDel(context.Background(), key, fields...).Err()
}

func (r *RedisSingle) HDel(key string, field string) error {
	key = RealKey(key)
	return r.client.HDel(context.Background(), key, field).Err()
}

func (r *RedisSingle) HMSet(key string, values map[string]string) error {
	key = RealKey(key)
	return r.client.HMSet(context.Background(), key, values).Err()
}

func (r *RedisSingle) CHMSet(key string, values map[string]interface{}, expire time.Duration) (err error) {
	var (
		pipe = r.client.Pipeline()
	)
	key = RealKey(key)
	pipe.HMSet(context.Background(), key, values)
	if expire > 0 {
		pipe.Expire(context.Background(), key, expire)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (r *RedisSingle) CBatchHSet(keys []string, field string, values []string) (err error) {
	if len(keys) != len(values) {
		return
	}
	var (
		i    int
		key  string
		pipe = r.client.Pipeline()
	)
	for i, key = range keys {
		pipe.HSet(context.Background(), RealKey(key), field, values[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (r *RedisSingle) HMGet(key string, fields ...string) []interface{} {
	key = RealKey(key)
	return r.client.HMGet(context.Background(), key, fields...).Val()
}

func (r *RedisSingle) HGet(key string, field string) (val string, err error) {
	key = RealKey(key)
	val, err = r.client.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisSingle) CHDel(keys []string, fields []string) (err error) {
	if len(keys) == 0 || len(fields) == 0 {
		return
	}
	if len(keys) != len(fields) {
		return
	}
	var (
		i    int
		key  string
		pipe = r.client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.HDel(context.Background(), key, fields[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

// Sequence ID
func (r *RedisSingle) GetMaxSeqID(chatId int64) (seqId uint64, err error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	seqId, err = GetUint64(key)
	if err == redis.Nil {
		err = nil
	}
	return
}

func (r *RedisSingle) IncrSeqID(chatId int64) (int64, error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	return Incr(key)
}

func (r *RedisSingle) DecrSeqID(chatId int64) (int64, error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	return Decr(key)
}

func (r *RedisSingle) SAdd(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return r.client.SAdd(context.Background(), key, members).Err()
}

func (r *RedisSingle) SRem(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return r.client.SRem(context.Background(), key, members).Err()
}

func (r *RedisSingle) SMembers(key string) []string {
	key = RealKey(key)
	return r.client.SMembers(context.Background(), key).Val()
}

func (r *RedisSingle) EvalSha(sha string, keys []string, args []interface{}) error {
	return r.client.EvalSha(context.Background(), sha, keys, args).Err()
}

func (r *RedisSingle) EvalShaResult(sha string, keys []string, args []interface{}) (interface{}, error) {
	return r.client.EvalSha(context.Background(), sha, keys, args).Result()
}

func (r *RedisSingle) ZAdd(key string, score float64, member string) (err error) {
	key = RealKey(key)
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	err = r.client.ZAdd(context.Background(), key, z).Err()
	return
}

func (r *RedisSingle) ZRem(key string, member string) (err error) {
	key = RealKey(key)
	err = r.client.ZRem(context.Background(), key, member).Err()
	return
}

func (r *RedisSingle) ZRevRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return r.client.ZRevRange(context.Background(), key, start, stop).Val()
}

func (r *RedisSingle) ZMScore(key string, members ...string) []float64 {
	key = RealKey(key)
	return r.client.ZMScore(context.Background(), key, members...).Val()
}

func (r *RedisSingle) ZRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return r.client.ZRange(context.Background(), key, start, stop).Val()
}

func (r *RedisSingle) ZRank(key, member string) (int64, error) {
	key = RealKey(key)
	return r.client.ZRank(context.Background(), key, member).Result()
}

func (r *RedisSingle) GeoAdd(key string, geoLocation ...*redis.GeoLocation) (err error) {
	key = RealKey(key)
	return r.client.GeoAdd(context.Background(), key, geoLocation...).Err()
}

func (r *RedisSingle) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return r.client.GeoRadius(context.Background(), key, longitude, latitude, query).Val()
}

func (r *RedisSingle) NewMutex(key string, options ...redsync.Option) *redsync.Mutex {
	key = r.RealKey(key)
	return r.redsSync.NewMutex(key, options...)
}

func (r *RedisSingle) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	key = r.RealKey(key)
	return r.client.SetNX(context.Background(), key, value, expiration).Result()
}

func (r *RedisSingle) ZIncrBy(key string, increment float64, member string) (float64, error) {
	key = r.RealKey(key)
	return r.client.ZIncrBy(context.Background(), key, increment, member).Result()
}
