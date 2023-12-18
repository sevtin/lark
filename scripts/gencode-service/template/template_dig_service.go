package template

var DigServiceTemplate = ParseTemplate(`
package dig

import (
	"lark/apps/apis/{{.PackageName}}/internal/service/svc_{{.ServiceName}}"
)

func init() {
	Provide(svc_{{.ServiceName}}.New{{.UpperServiceName}}Service)
}
`)
