package template

var DomainRepoTemplate = ParseTemplate(`
package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type {{.UpperModelName}}Repository interface {

}

type {{.LowerModelName}}Repository struct {
}

func New{{.UpperModelName}}Repository() {{.UpperModelName}}Repository {
	return &{{.LowerModelName}}Repository{}
}

func (r *{{.LowerModelName}}Repository) Create{{.UpperModelName}}({{.LowerModelName}} *po.{{.UpperModelName}}) (err error) {
	db:= xmysql.GetDB()
	err = db.Create({{.LowerModelName}}).Error
	return
}

func (r *{{.LowerModelName}}Repository) Update{{.UpperModelName}}(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = u.Updates(db, new(po.{{.UpperModelName}})).Error
	return
}

func (r *{{.LowerModelName}}Repository) {{.UpperModelName}}Info(q *entity.MysqlQuery)({{.LowerModelName}} *po.{{.UpperModelName}}, err error) {
	{{.LowerModelName}} = new(po.{{.UpperModelName}})
	db := xmysql.GetDB()
	q.Model = {{.LowerModelName}}
	err = q.Find(db, {{.LowerModelName}})
	return
}
`)
