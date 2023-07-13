package template

var DomainCacheTemplate = ParseTemplate(`
package cache

type {{.UpperServiceName}}Cache interface {
	
}

type {{.LowerServiceName}}Cache struct {
}

func New{{.UpperServiceName}}Cache() {{.UpperServiceName}}Cache {
	return &{{.LowerServiceName}}Cache{}
}
`)
