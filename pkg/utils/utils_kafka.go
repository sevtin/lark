package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func GetChatPartition[T constraints.Integer](num T) string {
	return fmt.Sprintf("_%v", num%10)
}
