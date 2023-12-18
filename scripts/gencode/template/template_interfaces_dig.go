package template

var InterfacesDigTemplate = ParseTemplate(`
package dig

import (
	"lark/apps/interfaces/internal/service/svc_{{.PackageName}}"
)

func init () {
	Provide(svc_{{.PackageName}}.New{{.UpperPackageName}}Service)
}
`)
