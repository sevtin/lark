package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	// 1、使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 2、跨域
	//engine.Use(middleware.Cors())
	return &GinServer{engine}
}

func (s *GinServer) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.Use(middleware...)
}

func (s *GinServer) Run(port int) {
	var (
		addr string
		err  error
	)
	//go func(p int) {
	//	log.Println(http.ListenAndServe(":"+strconv.Itoa(port+1), nil))
	//}(port)

	addr = ":" + strconv.Itoa(port)
	err = s.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed.", err.Error())
	}
}
