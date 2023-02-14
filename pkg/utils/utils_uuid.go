package utils

import (
	uuid "github.com/satori/go.uuid"
)

func NewUUID() (id string) {
	id = uuid.NewV4().String()
	return
}
