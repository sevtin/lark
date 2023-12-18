package dig

import (
	"lark/apps/interfaces/internal/service/svc_convo"
)

func init() {
	Provide(svc_convo.NewConvoService)
}
