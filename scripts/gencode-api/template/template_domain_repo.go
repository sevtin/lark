package template

var DomainRepoTemplate = ParseTemplate(`
package repo

type {{.UpperServiceName}}Repository interface {

}

type {{.LowerServiceName}}Repository struct {
}

func New{{.UpperServiceName}}Repository() {{.UpperServiceName}}Repository {
	return &{{.LowerServiceName}}Repository{}
}
`)
