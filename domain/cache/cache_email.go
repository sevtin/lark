package cache

type EmailCache interface {
}

type emailCache struct {
}

func NewEmailCache() EmailCache {
	return &emailCache{}
}
