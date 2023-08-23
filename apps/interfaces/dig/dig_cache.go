package dig

import (
	"lark/domain/cache"
)

func provideCache() {
	Provide(cache.NewServerMgrCache)
}
