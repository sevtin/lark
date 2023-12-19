package template

var DomainCrReadTemplate = ParseTemplate(`
package cr_{{.ModelName}}

import (
	"lark/domain/cache"
	"lark/domain/repo"
)

func Get{{.UpperModelName}}({{.LowerModelName}}Cache cache.{{.UpperModelName}}Cache, {{.LowerModelName}}Repo repo.{{.UpperModelName}}Repository, id int64) (err error) {
	return
}
`)
