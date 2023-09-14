package template

var DomainCrConstTemplate = ParseTemplate(`
package cr_{{.PackageName}}

import "errors"

var (
	ERROR_CR_CHAT_QUERY_FAILED = errors.New("查询失败")
)
`)
