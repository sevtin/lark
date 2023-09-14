package template

var InternalRouterTemplate = ParseTemplate(`
package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/{{.PackageName}}/dig"
	"lark/apps/apis/{{.PackageName}}/internal/ctrl"
	"lark/apps/apis/{{.PackageName}}/internal/service"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	group := engine.Group("api")
	group.Use(middleware.JwtAuth())
	register{{.UpperServiceName}}Router(group)
}

func register{{.UpperServiceName}}Router(group *gin.RouterGroup) {
	var svc service.{{.UpperServiceName}}Service
	dig.Invoke(func(s service.{{.UpperServiceName}}Service) {
		svc = s
	})
	ctrl := ctrl.New{{.UpperServiceName}}Ctrl(svc)
	router := group.Group("{{.PackageName}}")
	router.POST("edit", ctrl.Edit)
	router.GET("info", ctrl.Info)
}
`)
