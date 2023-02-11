package dig

import (
	"go.uber.org/dig"
	"lark/apps/upload/internal/config"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	provideUpload()
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
