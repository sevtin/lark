package xredis

import (
	"github.com/redis/go-redis/v9"
	"time"
)

func RealKey(key string) string {
	return cli.RealKey(key)
}

func GetPrefix() string {
	return cli.GetPrefix()
}

// 集群版用 最好别用
func GetClient() *redis.ClusterClient {
	return cli.GetClient()
}

func Pipeline() redis.Pipeliner {
	return cli.Pipeline()
}

func Unlink(key string) error {
	return cli.Unlink(key)
}

func TTL(key string) time.Duration {
	return cli.TTL(key)
}

func Del(key string) error {
	return cli.Del(key)
}

func CUnlink(keys []string) (err error) {
	return cli.CUnlink(keys)
}

func KeyExists(key string) (ok bool) {
	return cli.KeyExists(key)
}

func Set(key string, value interface{}, expire time.Duration) error {
	return cli.Set(key, value, expire)
}

func CSet(keys []string, values []interface{}, expire time.Duration) (err error) {
	return cli.CSet(keys, values, expire)
}

func CSets(keys []string, values []interface{}, expires []time.Duration) (err error) {
	return cli.CSets(keys, values, expires)
}

func Expire(key string, expire time.Duration) error {
	return cli.Expire(key, expire)
}

func Get(key string) (val string, err error) {
	return cli.Get(key)
}

func MGet(keys []string) ([]interface{}, error) {
	return cli.MGet(keys)
}

func CMGet(keys []string) (list []string, err error) {
	return cli.CMGet(keys)
}

func SlotMGet(maps map[uint16][]string) (list []interface{}, err error) {
	return cli.SlotMGet(maps)
}

func MSet(values ...interface{}) error {
	return cli.MSet(values)
}

func Incr(key string) (int64, error) {
	return cli.Incr(key)
}

func Decr(key string) (int64, error) {
	return cli.Decr(key)
}

func GetUint64(key string) (val uint64, err error) {
	return cli.GetUint64(key)
}

func GetInt(key string) (val int, err error) {
	return cli.GetInt(key)
}

func HGetInt64(key, field string) (value int64, err error) {
	return cli.HGetInt64(key, field)
}

func HGetAll(key string) map[string]string {
	return cli.HGetAll(key)
}

func HLen(key string) int64 {
	return cli.HLen(key)
}

func HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return cli.HScan(key, cursor, match, count)
}

func HSet(key string, value interface{}) error {
	return cli.HSet(key, value)
}

func HSetNX(key, field string, value interface{}) error {
	return cli.HSetNX(key, field, value)
}

func HSetNXEx(key, field string, value interface{}, ex time.Duration) (err error) {
	return cli.HSetNXEx(key, field, value, ex)
}

func HDels(key string, fields []string) error {
	return cli.HDels(key, fields)
}

func HDel(key string, field string) error {
	return cli.HDel(key, field)
}

func HMSet(key string, values map[string]string) error {
	return cli.HMSet(key, values)
}

func CHMSet(key string, values map[string]string, expire time.Duration) (err error) {
	return cli.CHMSet(key, values, expire)
}

func CBatchHSet(keys []string, field string, values []string) (err error) {
	return cli.CBatchHSet(keys, field, values)
}

func HMGet(key string, fields ...string) []interface{} {
	return cli.HMGet(key, fields...)
}

func HGet(key string, field string) (val string, err error) {
	return cli.HGet(key, field)
}

func CHDel(maps map[string][]string) (err error) {
	return cli.CHDel(maps)
}

// Sequence ID
func GetMaxSeqID(chatId int64) (seqId uint64, err error) {
	return cli.GetMaxSeqID(chatId)
}

func IncrSeqID(chatId int64) (int64, error) {
	return cli.IncrSeqID(chatId)
}

func DecrSeqID(chatId int64) (int64, error) {
	return cli.DecrSeqID(chatId)
}

func SAdd(key string, members ...interface{}) (err error) {
	return cli.SAdd(key, members...)
}

func SRem(key string, members ...interface{}) (err error) {
	return cli.SRem(key, members...)
}

func SMembers(key string) []string {
	return cli.SMembers(key)
}

func EvalSha(sha string, keys []string, args []interface{}) error {
	return cli.EvalSha(sha, keys, args)
}

func EvalShaResult(sha string, keys []string, args []interface{}) (interface{}, error) {
	return cli.EvalShaResult(sha, keys, args)
}

func ZAdd(key string, score float64, member string) (err error) {
	return cli.ZAdd(key, score, member)
}

func ZRem(key string, member string) (err error) {
	return cli.ZRem(key, member)
}

func ZRevRange(key string, start int64, stop int64) []string {
	return cli.ZRevRange(key, start, stop)
}

func ZMScore(key string, members ...string) []float64 {
	return cli.ZMScore(key, members...)
}

func ZRange(key string, start int64, stop int64) []string {
	return cli.ZRange(key, start, stop)
}

func ZRank(key, member string) (int64, error) {
	return cli.ZRank(key, member)
}

func GeoAdd(key string, geoLocation ...*redis.GeoLocation) (err error) {
	return cli.GeoAdd(key, geoLocation...)
}

func GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) []redis.GeoLocation {
	return cli.GeoRadius(key, longitude, latitude, query)
}

func ZIncrBy(key string, increment float64, member string) (float64, error) {
	return cli.ZIncrBy(key, increment, member)
}
