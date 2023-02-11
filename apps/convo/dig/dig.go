package dig

import (
	"go.uber.org/dig"
	"lark/apps/convo/internal/config"
	"lark/apps/convo/internal/server"
	"lark/apps/convo/internal/server/convo"
	"lark/apps/convo/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(convo.NewConvoServer)
	container.Provide(service.NewConvoService)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(cache.NewConvoCache)
	container.Provide(cache.NewChatMessageCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
