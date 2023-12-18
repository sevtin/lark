package xredis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisIface interface {
	Single() bool
	RealKey(key string) string
	GetPrefix() string
	GetClient() *redis.ClusterClient
	GetSingleClient() *redis.Client
	Pipeline() redis.Pipeliner
	Unlink(key string) error
	TTL(key string) time.Duration
	Del(key string) error
	CUnlink(keys []string) (err error)
	KeyExists(key string) (ok bool)
	Set(key string, value interface{}, expire time.Duration) error
	CSet(keys []string, values []interface{}, expire time.Duration) (err error)
	CSets(keys []string, values []interface{}, expires []time.Duration) (err error)
	Expire(key string, expire time.Duration) error
	Get(key string) (val string, err error)
	MGet(keys []string) ([]interface{}, error)
	CMGet(keys []string) (list []string, err error)
	SlotMGet(maps map[uint16][]string) (list []interface{}, err error)
	MSet(values ...interface{}) error
	Incr(key string) (int64, error)
	Decr(key string) (int64, error)
	GetUint64(key string) (val uint64, err error)
	GetInt(key string) (val int, err error)
	HGetInt64(key, field string) (value int64, err error)
	HGetAll(key string) map[string]string
	HLen(key string) int64
	HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)
	HSet(key string, value interface{}) error
	HSetNX(key, field string, value interface{}) error
	HDels(key string, fields []string) error
	HDel(key string, field string) error
	HMSet(key string, values map[string]string) error
	CHMSet(key string, values map[string]interface{}, expire time.Duration) (err error)
	CBatchHSet(keys []string, field string, values []string) (err error)
	HMGet(key string, fields ...string) []interface{}
	HGet(key string, field string) (val string, err error)
	CHDel(keys []string, fields []string) (err error)
	GetMaxSeqID(chatId int64) (seqId uint64, err error)
	IncrSeqID(chatId int64) (int64, error)
	DecrSeqID(chatId int64) (int64, error)
	SAdd(key string, members ...interface{}) (err error)
	SRem(key string, members ...interface{}) (err error)
	SMembers(key string) []string
	EvalSha(sha string, keys []string, args []interface{}) error
	EvalShaResult(sha string, keys []string, args []interface{}) (interface{}, error)
	ZAdd(key string, score float64, member string) (err error)
	ZRem(key string, member string) (err error)
	ZRevRange(key string, start int64, stop int64) []string
	ZMScore(key string, members ...string) []float64
	ZRange(key string, start int64, stop int64) []string
	ZRank(key, member string) (int64, error)
	GeoAdd(key string, geoLocation ...*redis.GeoLocation) (err error)
	GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation
	NewMutex(key string, options ...redsync.Option) *redsync.Mutex
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)
	ZIncrBy(key string, increment float64, member string) (float64, error)
}
