package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_hot/internal/config"
	"lark/apps/msg_hot/internal/server"
	"lark/apps/msg_hot/internal/server/msg_hot"
	"lark/apps/msg_hot/internal/service"
	"lark/domain/mrepo"
	"log/slog"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(mrepo.NewMessageHotRepository)
	Provide(msg_hot.NewMessageHotServer)
	Provide(service.NewMessageHotService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		slog.Warn(err.Error())
	}
}
