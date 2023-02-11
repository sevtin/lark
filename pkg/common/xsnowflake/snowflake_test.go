package xsnowflake

import (
	"fmt"
	"testing"
)

func TestGetSnowflakeID(t *testing.T) {
	id := NewSnowflakeID()
	fmt.Println(id)
}
