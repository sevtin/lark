package dig

import (
	"go.uber.org/dig"
	"lark/apps/interfaces/internal/config"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		log.Panic(err)
	}
}
