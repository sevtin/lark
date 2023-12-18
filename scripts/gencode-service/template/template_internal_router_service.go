package template

var InternalRouterServiceTemplate = ParseTemplate(`
package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/{{.PackageName}}/dig"
	"lark/apps/apis/{{.PackageName}}/internal/ctrl/ctrl_{{.ServiceName}}"
	"lark/apps/apis/{{.PackageName}}/internal/service/svc_{{.ServiceName}}"
)

func register{{.UpperServiceName}}Router(group *gin.RouterGroup) {
	var svc svc_{{.ServiceName}}.{{.UpperServiceName}}Service
	dig.Invoke(func(s svc_{{.ServiceName}}.{{.UpperServiceName}}Service) {
		svc = s
	})
	ctrl := ctrl_{{.ServiceName}}.New{{.UpperServiceName}}Ctrl(svc)
	router := group.Group("{{.ServiceName}}")
	router.POST("edit", ctrl.Edit)
	router.GET("info", ctrl.Info)
}
`)
