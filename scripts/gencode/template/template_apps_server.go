package template

var AppsServerTemplate = ParseTemplate(`
package server

import (
	"lark/apps/{{.PackageName}}/internal/server/{{.PackageName}}"
	"lark/pkg/commands"
)

type server struct {
	{{.LowerPackageName}}Server {{.PackageName}}.{{.UpperPackageName}}Server
}

func NewServer({{.LowerPackageName}}Server {{.PackageName}}.{{.UpperPackageName}}Server) commands.MainInstance {
	return &server{ {{.LowerPackageName}}Server: {{.LowerPackageName}}Server }
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.{{.LowerPackageName}}Server.Run()
}

func (s *server) Destroy() {

}
`)
