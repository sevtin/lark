package template

var AppsDigTemplate = ParseTemplate(`
package dig

import (
	"go.uber.org/dig"
	"lark/apps/{{.PackageName}}/internal/config"
	"lark/apps/{{.PackageName}}/internal/server"
	"lark/apps/{{.PackageName}}/internal/server/{{.PackageName}}"
	"lark/apps/{{.PackageName}}/internal/service"
	"log/slog"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide({{.PackageName}}.New{{.UpperPackageName}}Server)
	Provide(service.New{{.UpperPackageName}}Service)
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
