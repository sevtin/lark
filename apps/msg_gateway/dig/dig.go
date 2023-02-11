package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server"
	"lark/apps/msg_gateway/internal/server/gateway"
	"lark/apps/msg_gateway/internal/service"
	"lark/domain/cache"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(service.NewWsService)
	container.Provide(gateway.NewGatewayServer)
	container.Provide(cache.NewServerMgrCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
