package dig

import (
	"lark/apps/interfaces/internal/service/svc_payment"
)

func init() {
	Provide(svc_payment.NewPaymentService)
}
