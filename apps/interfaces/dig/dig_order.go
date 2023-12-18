package dig

import (
	"lark/apps/interfaces/internal/service/svc_order"
)

func init() {
	Provide(svc_order.NewOrderService)
}
