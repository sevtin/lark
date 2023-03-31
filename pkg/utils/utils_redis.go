package utils

import "fmt"

const (
	REDIS_CLUSTER_SLOT int64 = 16384
)

func GetHashTagKey(val int64) string {
	return fmt.Sprintf("{%d}", val)
}
