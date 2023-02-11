package dig

import (
	"lark/domain/cache"
)

func provideCache() {
	container.Provide(cache.NewServerMgrCache)
}
