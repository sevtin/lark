package template

var InterfacesDtoTemplate = ParseTemplate(`
package {{.PackageName}}

type {{.UpperServiceName}}EditReq struct {

}

type {{.UpperServiceName}}InfoReq struct {

}

`)
