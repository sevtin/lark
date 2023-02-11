package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetGetUUID(t *testing.T) {
	for {
		id := NewUUID()
		if strings.Contains(id, "2") == false {
			fmt.Println(id)
			break
		}
	}
}
