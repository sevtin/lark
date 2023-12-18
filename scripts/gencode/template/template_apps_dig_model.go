package template

var AppsDigModelTemplate = ParseTemplate(`
package dig

import (
	"lark/domain/cache"
	"lark/domain/repo"
)

func init() {
	Provide(repo.New{{.UpperModelName}}Repository)
	Provide(cache.New{{.UpperModelName}}Cache)
}

`)
