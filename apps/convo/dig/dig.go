package dig

import (
	"go.uber.org/dig"
	"lark/apps/convo/internal/config"
	"lark/apps/convo/internal/server"
	"lark/apps/convo/internal/server/convo"
	"lark/apps/convo/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(convo.NewConvoServer)
	Provide(service.NewConvoService)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewConvoCache)
	Provide(cache.NewChatMessageCache)
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
