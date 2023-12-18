package template

var InterfacesServiceApiTemplate = ParseTemplate(`
package svc_{{.PackageName}}

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_{{.PackageName}}"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_{{.PackageName}}"
	"lark/pkg/xhttp"
)

func (s *{{.LowerPackageName}}Service) {{.UpperApiName}}(params *dto_{{.PackageName}}.{{.UpperApiName}}Req) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	//var (
	//	req       = new(pb_{{.PackageName}}.{{.UpperApiName}}Req)
	//	reply     *pb_{{.PackageName}}.{{.UpperApiName}}Resp
	//	ack = new(dto_{{.PackageName}}.{{.UpperApiName}}Resp)
	//)
	//copier.Copy(req, params)
	//reply = s.{{.LowerPackageName}}Client.{{.UpperApiName}}(req)
	//if reply == nil {
	//	resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
	//	xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
	//	return
	//}
	//if reply.Code > 0 {
	//	resp.SetResult(reply.Code, reply.Msg)
	//	xlog.Warn(reply.Code, reply.Msg)
	//	return
	//}
	//copier.Copy(ack, reply)
	//resp.Data = ack
	return
}
`)
