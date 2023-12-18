package template

var InternalRouterTemplate = ParseTemplate(`
package router

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	//engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	publicGroup := engine.Group("open")
	registerOpenRoutes(publicGroup)

	privateGroup := engine.Group("api")
	registerPrivateRouter(privateGroup)
}

// 无需验证
func registerOpenRoutes(group *gin.RouterGroup) {

}

// 需要验证
func registerPrivateRouter(group *gin.RouterGroup) {
	group.Use(middleware.JwtAuth())
}

// 开放式api无需验证会话id
func registerPublicRouter(group *gin.RouterGroup) {
	router := group.Group("public")
	registerPublicTestRouter(router)
}

func registerPublicTestRouter(group *gin.RouterGroup) {

}
`)
