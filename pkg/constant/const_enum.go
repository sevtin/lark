package constant

// 订单状态
type OrderStatus int

const (
	OrderPending       OrderStatus = iota // 待付款
	OrderPaid                             // 已付款
	OrderCancelled                        // 已取消
	OrderRefunded                         // 已退款
	OrderInvalid                          // 失效
	OrderPaymentFailed                    // 订单支付失败
)

// 支付状态
type PaymentStatus int

const (
	PaymentPending   PaymentStatus = iota // 待支付
	PaymentPaid                           // 已支付
	PaymentCancelled                      // 取消支付
	PaymentRefunded                       // 已退款
	PaymentInvalid                        // 失效
	PaymentFailed                         // 支付失败
)
