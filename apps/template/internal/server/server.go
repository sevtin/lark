package server

import "lark/pkg/commands"

type server struct {
	srv interface{}
}

func NewServer(srv interface{}) commands.MainInstance {
	return &server{srv: srv}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {

}

func (s *server) Destroy() {

}
