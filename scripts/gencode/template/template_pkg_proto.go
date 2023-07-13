package template

var PkgProtoTemplate = ParseTemplate(`
syntax ="proto3";
package pb_{{.PackageName}};
option go_package = "./pb_{{.PackageName}};pb_{{.PackageName}}";
`)
