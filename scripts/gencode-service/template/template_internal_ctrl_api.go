package template

var InternalCtrlApiTemplate = ParseTemplate(`
package ctrl_{{.ServiceName}}

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/{{.PackageName}}/internal/dto/dto_{{.ServiceName}}"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *{{.UpperServiceName}}Ctrl) {{.UpperApiName}}(ctx *gin.Context) {
	var (
		params = new(dto_{{.ServiceName}}.{{.UpperApiName}}Req)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	// POST
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	// GET
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
	resp = ctrl.{{.LowerServiceName}}Service.{{.UpperApiName}}(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

`)
