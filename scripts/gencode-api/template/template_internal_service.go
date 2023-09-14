package template

var InternalServiceTemplate = ParseTemplate(`
package service

import (
	"lark/apps/apis/{{.PackageName}}/internal/config"
	"lark/apps/apis/{{.PackageName}}/internal/dto"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/xhttp"
)

type {{.UpperServiceName}}Service interface {
	Edit(req *dto.{{.UpperServiceName}}EditReq) (resp *xhttp.Resp)
	Info(req *dto.{{.UpperServiceName}}InfoReq) (resp *xhttp.Resp)
}

type {{.LowerServiceName}}Service struct {
	cfg            *config.Config
	{{.LowerServiceName}}Repo  repo.{{.UpperServiceName}}Repository
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

func (s *{{.LowerServiceName}}Service) Edit(req *dto.{{.UpperServiceName}}EditReq) (resp *xhttp.Resp) {
	resp = &xhttp.Resp{}
	return
}

func (s *{{.LowerServiceName}}Service) Info(req *dto.{{.UpperServiceName}}InfoReq) (resp *xhttp.Resp) {
	resp = &xhttp.Resp{}
	return
}
`)
