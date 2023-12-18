package template

var InternalDtoTemplate = ParseTemplate(`
package dto_{{.ServiceName}}

type {{.UpperServiceName}}EditReq struct {
	Uid int64
}

type {{.UpperServiceName}}InfoReq struct {
	Uid int64
}
`)
