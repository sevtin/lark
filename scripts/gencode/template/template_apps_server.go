package template

var AppsServerTemplate = ParseTemplate(`
package server

import (
	"lark/apps/{{.PackageName}}/internal/server/{{.PackageName}}"
	"lark/pkg/commands"
)

type server struct {
	{{.LowerServiceName}}Server {{.PackageName}}.{{.UpperServiceName}}Server
}

func NewServer({{.LowerServiceName}}Server {{.PackageName}}.{{.UpperServiceName}}Server) commands.MainInstance {
	return &server{ {{.LowerServiceName}}Server: {{.LowerServiceName}}Server }
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.{{.LowerServiceName}}Server.Run()
}

func (s *server) Destroy() {

}
`)
