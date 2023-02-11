package server

import (
	"flag"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/router"
	"lark/pkg/commands"
	"lark/pkg/common/xgin"
)

type server struct {
	ginServer *xgin.GinServer
	cfg       *config.Config
}

func NewServer() commands.MainInstance {
	return &server{cfg: config.NewConfig()}
}

func (s *server) Initialize() (err error) {
	flag.Parse()
	s.ginServer = xgin.NewGinServer()
	router.Register(s.ginServer.Engine)
	return
}

func (s *server) RunLoop() {
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
