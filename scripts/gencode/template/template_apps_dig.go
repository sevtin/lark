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
	"log/slog"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide({{.PackageName}}.New{{.UpperServiceName}}Server)
	Provide(service.New{{.UpperServiceName}}Service)
	Provide(repo.New{{.UpperServiceName}}Repository)
	Provide(cache.New{{.UpperServiceName}}Cache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		slog.Warn(err.Error())
	}
}
`)
