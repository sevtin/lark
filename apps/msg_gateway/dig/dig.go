package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server"
	"lark/apps/msg_gateway/internal/server/gateway"
	"lark/apps/msg_gateway/internal/service"
	"lark/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(service.NewWsService)
	Provide(gateway.NewGatewayServer)
	Provide(cache.NewServerMgrCache)
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
