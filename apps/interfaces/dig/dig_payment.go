package dig

import (
	"lark/apps/interfaces/internal/service/svc_payment"
)

func providePayment() {
	Provide(svc_payment.NewPaymentService)
}
