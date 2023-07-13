package template

var AppsServiceTemplate = ParseTemplate(`
package service

import (
	"lark/apps/{{.PackageName}}/internal/config"
	"lark/domain/cache"
	"lark/domain/repo"
)

type {{.UpperServiceName}}Service interface {
	
}

type {{.LowerServiceName}}Service struct {
	cfg            *config.Config
	{{.LowerServiceName}}Repo     repo.{{.UpperServiceName}}Repository
	{{.LowerServiceName}}Cache cache.{{.UpperServiceName}}Cache
}

func New{{.UpperServiceName}}Service(cfg *config.Config, 
	{{.LowerServiceName}}Repo     repo.{{.UpperServiceName}}Repository,
	{{.LowerServiceName}}Cache cache.{{.UpperServiceName}}Cache) {{.UpperServiceName}}Service {
	svc := &{{.LowerServiceName}}Service{cfg: cfg,
		{{.LowerServiceName}}Repo:  {{.LowerServiceName}}Repo,
		{{.LowerServiceName}}Cache: {{.LowerServiceName}}Cache,
	}
	return svc
}
`)
