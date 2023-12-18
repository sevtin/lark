package template

var DomainCacheTemplate = ParseTemplate(`
package cache

type {{.UpperModelName}}Cache interface {
	
}

type {{.LowerModelName}}Cache struct {
}

func New{{.UpperModelName}}Cache() {{.UpperModelName}}Cache {
	return &{{.LowerModelName}}Cache{}
}
`)
