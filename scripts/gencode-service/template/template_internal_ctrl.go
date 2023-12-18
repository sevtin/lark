package template

var InternalCtrlTemplate = ParseTemplate(`
package ctrl_{{.ServiceName}}

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/{{.PackageName}}/internal/dto/dto_{{.ServiceName}}"
	"lark/apps/apis/{{.PackageName}}/internal/service/svc_{{.ServiceName}}"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

type {{.UpperServiceName}}Ctrl struct {
	{{.LowerServiceName}}Service svc_{{.ServiceName}}.{{.UpperServiceName}}Service
}

func New{{.UpperServiceName}}Ctrl({{.LowerServiceName}}Service svc_{{.ServiceName}}.{{.UpperServiceName}}Service) *{{.UpperServiceName}}Ctrl {
	return &{{.UpperServiceName}}Ctrl{ {{.LowerServiceName}}Service: {{.LowerServiceName}}Service}
}

func (ctrl *{{.UpperServiceName}}Ctrl) Edit(ctx *gin.Context) {
	var (
		params = new(dto_{{.ServiceName}}.{{.UpperServiceName}}EditReq)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	params.Uid = uid
	resp = ctrl.{{.LowerServiceName}}Service.Edit(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

func (ctrl *{{.UpperServiceName}}Ctrl) Info(ctx *gin.Context) {
	var (
		params = new(dto_{{.ServiceName}}.{{.UpperServiceName}}InfoReq)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	params.Uid = uid
	resp = ctrl.{{.LowerServiceName}}Service.Info(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
`)
