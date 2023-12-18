package dig

import (
	"lark/domain/cache"
)

func init() {
	Provide(cache.NewServerMgrCache)
}
