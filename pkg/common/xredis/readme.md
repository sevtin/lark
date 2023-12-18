//func (r *RedisCluster)Dels(keys ...string) error {
//	return r.Client.Del(context.Background(), keys...).Err()
//}

//func (r *RedisCluster)CDel(keys []string) (err error) {
//	var (
//		pipe = r.Client.Pipeline()
//		key  string
//	)
//	for _, key = range keys {
//		pipe.Del(context.Background(), RealKey(key))
//	}
//	_, err = pipe.Exec(context.Background())
//	return
//}
