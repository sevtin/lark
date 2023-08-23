package template

var InterfacesDtoTemplate = ParseTemplate(`
package dto_{{.PackageName}}

type {{.UpperServiceName}}EditReq struct {

}

type {{.UpperServiceName}}InfoReq struct {

}

`)
