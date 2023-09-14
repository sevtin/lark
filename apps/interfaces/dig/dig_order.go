package dig

import (
	"lark/apps/interfaces/internal/service/svc_order"
)

func provideOrder() {
	Provide(svc_order.NewOrderService)
}
