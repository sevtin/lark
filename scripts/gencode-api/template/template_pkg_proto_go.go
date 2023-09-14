package template

var PkgProtoGoTemplate = ParseTemplate(`
package pb_{{.PackageName}}

type Unimplemented{{.UpperServiceName}}Server struct {

}

func Register{{.UpperServiceName}}Server(s interface{}, srv interface{}) {

}
`)
