package template

var DomainRepoTemplate = ParseTemplate(`
package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
)

type {{.UpperModelName}}Repository interface {

}

type {{.LowerModelName}}Repository struct {
}

func New{{.UpperModelName}}Repository() {{.UpperModelName}}Repository {
	return &{{.LowerModelName}}Repository{}
}

func (r *{{.LowerModelName}}Repository) Create({{.LowerModelName}} *po.{{.UpperModelName}}) (err error) {
	db:= xmysql.GetDB()
	err = db.Create({{.LowerModelName}}).Error
	return
}
`)
