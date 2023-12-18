package template

var AppsServiceApiTemplate = ParseTemplate(`
package service

import (
	"context"
	"lark/pkg/proto/pb_{{.PackageName}}"
)

func (s *{{.LowerPackageName}}Service) {{.UpperApiName}}(ctx context.Context, req *pb_{{.PackageName}}.{{.UpperApiName}}Req) (resp *pb_{{.PackageName}}.{{.UpperApiName}}, _ error) {
	resp = new(pb_{{.PackageName}}.{{.UpperApiName}})
	return
}
`)
