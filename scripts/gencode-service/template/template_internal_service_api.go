package template

var InternalServiceApiTemplate = ParseTemplate(`
package svc_{{.ServiceName}}

import (
	"lark/apps/apis/{{.PackageName}}/internal/dto/dto_{{.ServiceName}}"
	"lark/pkg/xhttp"
)

func (s *{{.LowerServiceName}}Service) {{.UpperApiName}}(req *dto_{{.ServiceName}}.{{.UpperApiName}}Req) (resp *xhttp.Resp) {
	resp = &xhttp.Resp{}
	return
}
`)
