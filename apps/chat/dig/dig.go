package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat/internal/config"
	"lark/apps/chat/internal/server"
	"lark/apps/chat/internal/server/chat"
	"lark/apps/chat/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(chat.NewChatServer)
	container.Provide(service.NewChatService)
	container.Provide(repo.NewChatRepository)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(repo.NewChatInviteRepository)
	container.Provide(repo.NewUserRepository)
	container.Provide(repo.NewAvatarRepository)
	container.Provide(cache.NewChatCache)
	container.Provide(cache.NewChatMessageCache)
	container.Provide(cache.NewChatMemberCache)
	container.Provide(cache.NewUserCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
