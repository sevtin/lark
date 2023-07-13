package template

var InterfacesServiceTemplate = ParseTemplate(`
package svc_{{.PackageName}}

import (
	{{.PackageName}}_client "lark/apps/{{.PackageName}}/client"
	"lark/apps/interfaces/internal/config"
)

type {{.UpperServiceName}}Service interface {

}

type {{.LowerServiceName}}Service struct {
	{{.LowerServiceName}}Client {{.PackageName}}_client.{{.UpperServiceName}}Client
}

func New{{.UpperServiceName}}Service(conf *config.Config) {{.UpperServiceName}}Service {
	{{.LowerServiceName}}Client := {{.PackageName}}_client.New{{.UpperServiceName}}Client(conf.Etcd, nil, conf.Jaeger, conf.Name)
	return &{{.LowerServiceName}}Service{ {{.LowerServiceName}}Client: {{.LowerServiceName}}Client }
}
`)
