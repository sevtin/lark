package template

var InterfacesRouterTemplate = ParseTemplate(`
package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_{{.PackageName}}"
	"lark/apps/interfaces/internal/service/svc_{{.PackageName}}"
)

func register{{.UpperServiceName}}Router(group *gin.RouterGroup) {
	var svc svc_{{.PackageName}}.{{.UpperServiceName}}Service
	dig.Invoke(func(s svc_{{.PackageName}}.{{.UpperServiceName}}Service) {
		svc = s
	})
	ctrl := ctrl_{{.PackageName}}.New{{.UpperServiceName}}Ctrl(svc)
	router := group.Group("{{.PackageName}}")
	router.POST("edit", ctrl.Edit)
	router.GET("info", ctrl.Info)
}
`)
