package cache

type PaymentCache interface {
}

type paymentCache struct {
}

func NewPaymentCache() PaymentCache {
	return &paymentCache{}
}
