package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat_invite/internal/config"
	"lark/apps/chat_invite/internal/server"
	"lark/apps/chat_invite/internal/server/chat_invite"
	"lark/apps/chat_invite/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(chat_invite.NewChatInviteServer)
	container.Provide(service.NewChatInviteService)
	container.Provide(repo.NewChatInviteRepository)
	container.Provide(repo.NewUserRepository)
	container.Provide(repo.NewAvatarRepository)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(repo.NewChatRepository)
	container.Provide(cache.NewChatCache)
	container.Provide(cache.NewChatMessageCache)
	container.Provide(cache.NewChatMemberCache)
	container.Provide(cache.NewUserCache)

}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
