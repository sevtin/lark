package utils

import "fmt"

const (
	REDIS_CLUSTER_SLOT int64 = 16384
)

func GetHashTagKey(val interface{}) string {
	return fmt.Sprintf("{%v}", val)
}
