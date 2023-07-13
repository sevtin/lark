package template

var InterfacesCtrlTemplate = ParseTemplate(`
package ctrl_{{.PackageName}}

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/service/svc_{{.PackageName}}"
	"lark/pkg/xhttp"
)

type {{.UpperServiceName}}Ctrl struct {
	{{.LowerServiceName}}Service svc_{{.PackageName}}.{{.UpperServiceName}}Service
}

func New{{.UpperServiceName}}Ctrl({{.LowerServiceName}}Service svc_{{.PackageName}}.{{.UpperServiceName}}Service) *{{.UpperServiceName}}Ctrl {
	return &{{.UpperServiceName}}Ctrl{ {{.LowerServiceName}}Service: {{.LowerServiceName}}Service }
}

func (ctrl *{{.UpperServiceName}}Ctrl) Edit(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}

func (ctrl *{{.UpperServiceName}}Ctrl) Info(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}
`)
