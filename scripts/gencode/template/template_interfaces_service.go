package template

var InterfacesServiceTemplate = ParseTemplate(`
package svc_{{.PackageName}}

import (
	{{.PackageName}}_client "lark/apps/{{.PackageName}}/client"
	"lark/apps/interfaces/internal/config"
)

type {{.UpperPackageName}}Service interface {

}

type {{.LowerPackageName}}Service struct {
	{{.LowerPackageName}}Client {{.PackageName}}_client.{{.UpperPackageName}}Client
}

func New{{.UpperPackageName}}Service(conf *config.Config) {{.UpperPackageName}}Service {
	{{.LowerPackageName}}Client := {{.PackageName}}_client.New{{.UpperPackageName}}Client(conf.Etcd, nil, conf.Jaeger, conf.Name)
	return &{{.LowerPackageName}}Service{ {{.LowerPackageName}}Client: {{.LowerPackageName}}Client }
}
`)
