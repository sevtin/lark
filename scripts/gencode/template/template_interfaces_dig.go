package template

var InterfacesDigTemplate = ParseTemplate(`
package dig

import (
	"lark/apps/interfaces/internal/service/svc_{{.PackageName}}"
)

func provide{{.UpperServiceName}}() {
	Provide(svc_{{.PackageName}}.New{{.UpperServiceName}}Service)
}

`)
