package template

var InternalDtoTemplate = ParseTemplate(`
package dto

type {{.UpperServiceName}}EditReq struct {

}

type {{.UpperServiceName}}InfoReq struct {

}
`)
