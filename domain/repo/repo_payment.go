package repo

type PaymentRepository interface {
}

type paymentRepository struct {
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{}
}
