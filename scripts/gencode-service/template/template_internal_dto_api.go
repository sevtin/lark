package template

var InterfacesDtoApiTemplate = ParseTemplate(`
package dto_{{.ServiceName}}

type {{.UpperApiName}}Req struct {
	Uid int64
}

type {{.UpperApiName}}Resp struct {
	Uid int64
}
`)
