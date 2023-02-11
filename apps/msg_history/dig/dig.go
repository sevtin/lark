package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_history/internal/config"
	"lark/apps/msg_history/internal/server"
	"lark/apps/msg_history/internal/server/msg_history"
	"lark/apps/msg_history/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(msg_history.NewMessageHistoryServer)
	container.Provide(service.NewMessageHistoryService)
	container.Provide(repo.NewChatMessageRepository)
	container.Provide(cache.NewChatMessageCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
