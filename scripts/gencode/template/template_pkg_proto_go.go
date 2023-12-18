package template

var PkgProtoGoTemplate = ParseTemplate(`
package pb_{{.PackageName}}

type Unimplemented{{.UpperPackageName}}Server struct {

}

func Register{{.UpperPackageName}}Server(s interface{}, srv interface{}) {

}
`)
