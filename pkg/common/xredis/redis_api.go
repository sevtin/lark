package xredis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

func RealKey(key string) string {
	if cli != nil {
		return cli.Prefix + key
	}
	return key
}

func Del(key string) error {
	key = RealKey(key)
	return cli.client.Del(context.Background(), key).Err()
}

func TTL(key string) time.Duration {
	key = RealKey(key)
	return cli.client.TTL(context.Background(), key).Val()
}

func Dels(keys ...string) error {
	return cli.client.Del(context.Background(), keys...).Err()
}

func KeyExists(key string) (ok bool) {
	key = RealKey(key)
	val := cli.client.Exists(context.Background(), key).Val()
	if val == 1 {
		ok = true
	}
	return
}

func Set(key string, value interface{}, expire time.Duration) error {
	key = RealKey(key)
	if expire > 0 {
		return cli.client.Set(context.Background(), key, value, expire).Err()
	}
	return cli.client.Set(context.Background(), key, value, 0).Err()
}

func Expire(key string, expire time.Duration) error {
	key = RealKey(key)
	return cli.client.Expire(context.Background(), key, expire).Err()
}

func Get(key string) (val string, err error) {
	key = RealKey(key)
	val, err = cli.client.Get(context.Background(), key).Result()
	//if err == redis.Nil {
	//	err = nil
	//}
	return
}

func MGet(keys ...string) ([]interface{}, error) {
	return cli.client.MGet(context.Background(), keys...).Result()
}

func MSet(values ...interface{}) error {
	// MSET 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。
	return cli.client.MSet(context.Background(), values...).Err()
}

func Incr(key string) (int64, error) {
	key = RealKey(key)
	return cli.client.Incr(context.Background(), key).Result()
}

func Decr(key string) (int64, error) {
	key = RealKey(key)
	return cli.client.Decr(context.Background(), key).Result()
}

func GetUint64(key string) (uint64, error) {
	key = RealKey(key)
	return cli.client.Get(context.Background(), key).Uint64()
}

func GetInt(key string) (int, error) {
	key = RealKey(key)
	return cli.client.Get(context.Background(), key).Int()
}

func HGetInt64(key, field string) (value int64, err error) {
	key = RealKey(key)
	return cli.client.HGet(context.Background(), key, field).Int64()
}

func HGetAll(key string) map[string]string {
	key = RealKey(key)
	hash := cli.client.HGetAll(context.Background(), key).Val()
	return hash
}

func HSet(key string, value interface{}) error {
	key = RealKey(key)
	return cli.client.HSet(context.Background(), key, value).Err()
}

func HSetNX(key, field string, value interface{}) error {
	key = RealKey(key)
	return cli.client.HSetNX(context.Background(), key, field, value).Err()
}

func HDels(key string, fields []string) error {
	key = RealKey(key)
	return cli.client.HDel(context.Background(), key, fields...).Err()
}

func HDel(key string, field string) error {
	key = RealKey(key)
	return cli.client.HDel(context.Background(), key, field).Err()
}

func HMSet(key string, values map[string]interface{}) error {
	key = RealKey(key)
	return cli.client.HMSet(context.Background(), key, values).Err()
}

func HMGet(key string, fields ...string) []interface{} {
	key = RealKey(key)
	return cli.client.HMGet(context.Background(), key, fields...).Val()
}

func HGet(key string, field string) string {
	key = RealKey(key)
	return cli.client.HGet(context.Background(), key, field).Val()
}

// Sequence ID
func GetMaxSeqID(chatId int64) (seqId uint64, err error) {
	key := MSG_SEQ_ID + strconv.FormatInt(chatId, 10)
	seqId, err = GetUint64(key)
	if err == redis.Nil {
		err = nil
	}
	return
}

func IncrSeqID(chatId int64) (int64, error) {
	key := MSG_SEQ_ID + strconv.FormatInt(chatId, 10)
	return Incr(key)
}

func DecrSeqID(chatId int64) (int64, error) {
	key := MSG_SEQ_ID + strconv.FormatInt(chatId, 10)
	return Decr(key)
}

func SAdd(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return cli.client.SAdd(context.Background(), key, members).Err()
}

func SRem(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return cli.client.SRem(context.Background(), key, members).Err()
}

func SMembers(key string) []string {
	key = RealKey(key)
	return cli.client.SMembers(context.Background(), key).Val()
}

func EvalSha(sha string, keys []string, args []interface{}) error {
	return cli.client.EvalSha(context.Background(), sha, keys, args).Err()
}

func EvalShaResult(sha string, keys []string, args []interface{}) (interface{}, error) {
	return cli.client.EvalSha(context.Background(), sha, keys, args).Result()
}

// 可能只删除部分
func DelKeysByMatch(match string, timeout time.Duration) (err error) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		iter   *redis.ScanIterator
	)
	match = RealKey(match)
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter = cli.client.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err = cli.client.Del(ctx, iter.Val()).Err()
		if err != nil {
			return
		}
	}
	if err = iter.Err(); err != nil {
		return
	}
	return
}

func ZAdd(key string, score float64, member string) (err error) {
	key = RealKey(key)
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	err = cli.client.ZAdd(context.Background(), key, z).Err()
	return
}

func ZRem(key string, member string) (err error) {
	key = RealKey(key)
	err = cli.client.ZRem(context.Background(), key, member).Err()
	return
}

func ZRevRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return cli.client.ZRevRange(context.Background(), key, start, stop).Val()
}

func ZMScore(key string, members ...string) []float64 {
	key = RealKey(key)
	return cli.client.ZMScore(context.Background(), key, members...).Val()
}

func ZRange(key string, start int64, stop int64) []string {
	key = RealKey(key)
	return cli.client.ZRange(context.Background(), key, start, stop).Val()
}

func ZRank(key, member string) (int64, error) {
	key = RealKey(key)
	return cli.client.ZRank(context.Background(), key, member).Result()
}
