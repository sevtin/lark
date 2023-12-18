package template

var InterfacesDtoTemplate = ParseTemplate(`
package dto_{{.PackageName}}

type {{.UpperApiName}}Req struct {
	Uid int64
}

type {{.UpperApiName}}Resp struct {
	Uid int64
}
`)
