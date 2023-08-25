package xredis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"time"
)

func RealKey(key string) string {
	if Cli != nil {
		return Cli.Prefix + key
	}
	return key
}

func GetPrefix() string {
	if Cli != nil {
		return Cli.Prefix
	}
	return ""
}

func (r *RedisClient) RealKey(key string) string {
	return r.Prefix + key
}

//func Del(key string) error {
//	key = RealKey(key)
//	return Cli.Client.Del(context.Background(), key).Err()
//}

func Unlink(key string) error {
	key = RealKey(key)
	return Cli.Client.Unlink(context.Background(), key).Err()
}

func TTL(key string) time.Duration {
	key = RealKey(key)
	return Cli.Client.TTL(context.Background(), key).Val()
}

func Del(key string) error {
	key = RealKey(key)
	return Cli.Client.Del(context.Background(), key).Err()
}

//func Dels(keys ...string) error {
//	return Cli.Client.Del(context.Background(), keys...).Err()
//}

//func CDel(keys []string) (err error) {
//	var (
//		pipe = Cli.Client.Pipeline()
//		key  string
//	)
//	for _, key = range keys {
//		pipe.Del(context.Background(), RealKey(key))
//	}
//	_, err = pipe.Exec(context.Background())
//	return
//}

func CUnlink(keys []string) (err error) {
	var (
		pipe = Cli.Client.Pipeline()
		key  string
	)
	for _, key = range keys {
		pipe.Unlink(context.Background(), RealKey(key))
	}
	_, err = pipe.Exec(context.Background())
	return
}

func KeyExists(key string) (ok bool) {
	key = RealKey(key)
	val := Cli.Client.Exists(context.Background(), key).Val()
	if val == 1 {
		ok = true
	}
	return
}

func Set(key string, value interface{}, expire time.Duration) error {
	key = RealKey(key)
	if expire > 0 {
		return Cli.Client.Set(context.Background(), key, value, expire).Err()
	}
	return Cli.Client.Set(context.Background(), key, value, 0).Err()
}

func CSet(keys []string, values []interface{}, expire time.Duration) (err error) {
	if len(keys) != len(values) {
		return
	}
	var (
		i    int
		key  string
		pipe = Cli.Client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.Set(context.Background(), key, values[i], expire)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func CSets(keys []string, values []interface{}, expires []time.Duration) (err error) {
	if len(keys) != len(values) || len(values) != len(expires) {
		return
	}
	var (
		i    int
		key  string
		pipe = Cli.Client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.Set(context.Background(), key, values[i], expires[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

func Expire(key string, expire time.Duration) error {
	key = RealKey(key)
	return Cli.Client.Expire(context.Background(), key, expire).Err()
}

func Get(key string) (val string, err error) {
	key = RealKey(key)
	val, err = Cli.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func MGet(keys []string) ([]interface{}, error) {
	return Cli.Client.MGet(context.Background(), keys...).Result()
}

func CMGet(keys []string) (list []string, err error) {
	var (
		key     string
		cmdList = make([]*redis.StringCmd, len(keys))
		i       int
		c       *redis.StringCmd
		pipe    = Cli.Client.Pipeline()
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

func SlotMGet(maps map[uint16][]string) (list []interface{}, err error) {
	var (
		index   = -1
		keys    []string
		cmdList = make([]*redis.SliceCmd, len(maps))
		c       *redis.SliceCmd
		pipe    = Cli.Client.Pipeline()
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

func MSet(values ...interface{}) error {
	// MSET 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。
	return Cli.Client.MSet(context.Background(), values...).Err()
}

func Incr(key string) (int64, error) {
	key = RealKey(key)
	return Cli.Client.Incr(context.Background(), key).Result()
}

func Decr(key string) (int64, error) {
	key = RealKey(key)
	return Cli.Client.Decr(context.Background(), key).Result()
}

func GetUint64(key string) (val uint64, err error) {
	key = RealKey(key)
	val, err = Cli.Client.Get(context.Background(), key).Uint64()
	if err == redis.Nil {
		err = nil
	}
	return
}

func GetInt(key string) (int, error) {
	key = RealKey(key)
	return Cli.Client.Get(context.Background(), key).Int()
}

func HGetInt64(key, field string) (value int64, err error) {
	key = RealKey(key)
	return Cli.Client.HGet(context.Background(), key, field).Int64()
}

func HGetAll(key string) map[string]string {
	key = RealKey(key)
	hash := Cli.Client.HGetAll(context.Background(), key).Val()
	return hash
}

func HLen(key string) int64 {
	key = RealKey(key)
	return Cli.Client.HLen(context.Background(), key).Val()
}

func HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	key = RealKey(key)
	return Cli.Client.HScan(context.Background(), key, cursor, match, count).Result()
}

func HSet(key string, value interface{}) error {
	key = RealKey(key)
	return Cli.Client.HSet(context.Background(), key, value).Err()
}

func HSetNX(key, field string, value interface{}) error {
	key = RealKey(key)
	return Cli.Client.HSetNX(context.Background(), key, field, value).Err()
}

func HDels(key string, fields []string) error {
	key = RealKey(key)
	return Cli.Client.HDel(context.Background(), key, fields...).Err()
}

func HDel(key string, field string) error {
	key = RealKey(key)
	return Cli.Client.HDel(context.Background(), key, field).Err()
}

func HMSet(key string, values map[string]string) error {
	key = RealKey(key)
	return Cli.Client.HMSet(context.Background(), key, values).Err()
}

func CHMSet(key string, values map[string]interface{}, expire time.Duration) (err error) {
	var (
		pipe = Cli.Client.Pipeline()
	)
	key = RealKey(key)
	pipe.HMSet(context.Background(), key, values)
	if expire > 0 {
		pipe.Expire(context.Background(), key, expire)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func CBatchHSet(keys []string, field string, values []string) (err error) {
	if len(keys) != len(values) {
		return
	}
	var (
		i    int
		key  string
		pipe = Cli.Client.Pipeline()
	)
	for i, key = range keys {
		pipe.HSet(context.Background(), RealKey(key), field, values[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

func HMGet(key string, fields ...string) []interface{} {
	key = RealKey(key)
	return Cli.Client.HMGet(context.Background(), key, fields...).Val()
}

func HGet(key string, field string) (val string, err error) {
	key = RealKey(key)
	val, err = Cli.Client.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func CHDel(keys []string, fields []string) (err error) {
	if len(keys) == 0 || len(fields) == 0 {
		return
	}
	if len(keys) != len(fields) {
		return
	}
	var (
		i    int
		key  string
		pipe = Cli.Client.Pipeline()
	)
	for i, key = range keys {
		key = RealKey(key)
		pipe.HDel(context.Background(), key, fields[i])
	}
	_, err = pipe.Exec(context.Background())
	return
}

// Sequence ID
func GetMaxSeqID(chatId int64) (seqId uint64, err error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	seqId, err = GetUint64(key)
	if err == redis.Nil {
		err = nil
	}
	return
}

func IncrSeqID(chatId int64) (int64, error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	return Incr(key)
}

func DecrSeqID(chatId int64) (int64, error) {
	key := constant.RK_MSG_SEQ_ID + utils.GetHashTagKey(chatId)
	return Decr(key)
}

func SAdd(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return Cli.Client.SAdd(context.Background(), key, members).Err()
}

func SRem(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return Cli.Client.SRem(context.Background(), key, members).Err()
}

func SMembers(key string) []string {
	key = RealKey(key)
	return Cli.Client.SMembers(context.Background(), key).Val()
}

func EvalSha(sha string, keys []string, args []interface{}) error {
	return Cli.Client.EvalSha(context.Background(), sha, keys, args).Err()
}

func EvalShaResult(sha string, keys []string, args []interface{}) (interface{}, error) {
	return Cli.Client.EvalSha(context.Background(), sha, keys, args).Result()
}

// 可能只删除部分
//func DelKeysByMatch(match string, timeout time.Duration) (err error) {
//	var (
//		ctx    context.Context
//		cancel context.CancelFunc
//		iter   *redis.ScanIterator
//	)
//	match = RealKey(match)
//	ctx, cancel = context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	iter = Cli.Client.Scan(ctx, 0, match, 0).Iterator()
//	for iter.Next(ctx) {
//		err = Cli.Client.Del(ctx, iter.Val()).Err()
//		if err != nil {
//			return
//		}
//	}
//	if err = iter.Err(); err != nil {
//		return
//	}
//	return
//}

func ZAdd(key string, score float64, member string) (err error) {
	key = RealKey(key)
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	err = Cli.Client.ZAdd(context.Background(), key, z).Err()
	return
}

func ZRem(key string, member string) (err error) {
	key = RealKey(key)
	err = Cli.Client.ZRem(context.Background(), key, member).Err()
	return
}

func ZRevRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return Cli.Client.ZRevRange(context.Background(), key, start, stop).Val()
}

func ZMScore(key string, members ...string) []float64 {
	key = RealKey(key)
	return Cli.Client.ZMScore(context.Background(), key, members...).Val()
}

func ZRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return Cli.Client.ZRange(context.Background(), key, start, stop).Val()
}

func ZRank(key, member string) (int64, error) {
	key = RealKey(key)
	return Cli.Client.ZRank(context.Background(), key, member).Result()
}
