package template

var AppsServiceTemplate = ParseTemplate(`
package service

import (
	"lark/apps/{{.PackageName}}/internal/config"
)

type {{.UpperPackageName}}Service interface {
	
}

type {{.LowerPackageName}}Service struct {
	cfg *config.Config
}

func New{{.UpperPackageName}}Service(cfg *config.Config) {{.UpperPackageName}}Service {
	svc := &{{.LowerPackageName}}Service{cfg: cfg}
	return svc
}
`)
