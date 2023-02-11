package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func NewUUID() (id string) {
	id = fmt.Sprintf("%s", uuid.NewV4())
	return
}
