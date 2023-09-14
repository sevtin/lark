package template

var DomainCrReadTemplate = ParseTemplate(`
package cr_{{.PackageName}}

import (
	"lark/domain/cache"
	"lark/domain/repo"
)

func Get{{.UpperServiceName}}Info({{.LowerServiceName}}Cache cache.{{.UpperServiceName}}Cache, {{.LowerServiceName}}Repo repo.{{.UpperServiceName}}Repository, id int64) (err error) {
	return
}
`)
