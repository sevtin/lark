package server

import (
	"flag"
	"lark/apps/apis/upload/internal/config"
	"lark/apps/apis/upload/internal/router"
	"lark/pkg/commands"
	"lark/pkg/common/xetcd"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
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
	err := xetcd.RegisterEtcd(s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, utils.GetServerIP(), s.cfg.Port, s.cfg.Name, xetcd.TIME_TO_LIVE)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
