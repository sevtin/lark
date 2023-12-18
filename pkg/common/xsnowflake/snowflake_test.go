package xsnowflake

import (
	"fmt"
	"testing"
)

func TestGetSnowflakeID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := NewSnowflakeID()
		fmt.Println(id)
	}
}
