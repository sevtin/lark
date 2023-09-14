package cache

type OrderCache interface {
}

type orderCache struct {
}

func NewOrderCache() OrderCache {
	return &orderCache{}
}
