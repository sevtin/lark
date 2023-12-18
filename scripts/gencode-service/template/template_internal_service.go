package template

var InternalServiceTemplate = ParseTemplate(`
package svc_{{.ServiceName}}

import (
	"lark/apps/apis/{{.PackageName}}/internal/config"
	"lark/apps/apis/{{.PackageName}}/internal/dto/dto_{{.ServiceName}}"
	"lark/pkg/xhttp"
)

type {{.UpperServiceName}}Service interface {
	Edit(req *dto_{{.ServiceName}}.{{.UpperServiceName}}EditReq) (resp *xhttp.Resp)
	Info(req *dto_{{.ServiceName}}.{{.UpperServiceName}}InfoReq) (resp *xhttp.Resp)
}

type {{.LowerServiceName}}Service struct {
	cfg            *config.Config
}

func New{{.UpperServiceName}}Service(cfg *config.Config) {{.UpperServiceName}}Service {
	svc := &{{.LowerServiceName}}Service{cfg: cfg}
	return svc
}

func (s *{{.LowerServiceName}}Service) Edit(req *dto_{{.ServiceName}}.{{.UpperServiceName}}EditReq) (resp *xhttp.Resp) {
	resp = &xhttp.Resp{}
	return
}

func (s *{{.LowerServiceName}}Service) Info(req *dto_{{.ServiceName}}.{{.UpperServiceName}}InfoReq) (resp *xhttp.Resp) {
	resp = &xhttp.Resp{}
	return
}
`)
