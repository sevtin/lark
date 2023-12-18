package template

var DomainCrConstTemplate = ParseTemplate(`
package cr_{{.ModelName}}

import "errors"

var (
	ERROR_CR_{{.AllUpperModelName}}_QUERY_FAILED = errors.New("查询失败")
)
`)
