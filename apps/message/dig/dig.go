package dig

import (
	"go.uber.org/dig"
	"lark/apps/message/internal/config"
	"lark/apps/message/internal/server"
	"lark/apps/message/internal/server/message"
	"lark/apps/message/internal/service"
	"lark/domain/cache"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(message.NewMessageServer)
	container.Provide(service.NewMessageService)
	container.Provide(cache.NewChatMemberCache)
	container.Provide(cache.NewChatMessageCache)

}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
