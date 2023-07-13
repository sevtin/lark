package template

var AppsDigTemplate = ParseTemplate(`
package dig

import (
	"go.uber.org/dig"
	"lark/apps/{{.PackageName}}/internal/config"
	"lark/apps/{{.PackageName}}/internal/server"
	"lark/apps/{{.PackageName}}/internal/server/{{.PackageName}}"
	"lark/apps/{{.PackageName}}/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide({{.PackageName}}.New{{.UpperServiceName}}Server)
	container.Provide(service.New{{.UpperServiceName}}Service)
	container.Provide(repo.New{{.UpperServiceName}}Repository)
	container.Provide(cache.New{{.UpperServiceName}}Cache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
`)
