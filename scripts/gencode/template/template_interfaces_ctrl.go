package template

var InterfacesCtrlTemplate = ParseTemplate(`
package ctrl_{{.PackageName}}

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/service/svc_{{.PackageName}}"
	"lark/pkg/xhttp"
)

type {{.UpperPackageName}}Ctrl struct {
	{{.LowerPackageName}}Service svc_{{.PackageName}}.{{.UpperPackageName}}Service
}

func New{{.UpperPackageName}}Ctrl({{.LowerPackageName}}Service svc_{{.PackageName}}.{{.UpperPackageName}}Service) *{{.UpperPackageName}}Ctrl {
	return &{{.UpperPackageName}}Ctrl{ {{.LowerPackageName}}Service: {{.LowerPackageName}}Service }
}

func (ctrl *{{.UpperPackageName}}Ctrl) Edit(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}

func (ctrl *{{.UpperPackageName}}Ctrl) Info(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}
`)
