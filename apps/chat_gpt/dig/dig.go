package dig

import (
	"go.uber.org/dig"
)

var container = dig.New()

func init() {
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
