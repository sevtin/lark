package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat_msg/internal/config"
	"lark/apps/chat_msg/internal/server"
	"lark/apps/chat_msg/internal/server/chat_msg"
	"lark/apps/chat_msg/internal/service"
	"lark/domain/cache"
	"lark/domain/mrepo"
	"lark/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat_msg.NewChatMessageServer)
	Provide(service.NewChatMessageService)
	Provide(repo.NewChatMessageRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(mrepo.NewMessageHotRepository)
	Provide(cache.NewChatMessageCache)
	Provide(cache.NewChatMemberCache)
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
