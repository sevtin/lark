package dig

import (
	"lark/apps/interfaces/internal/service/svc_convo"
)

func provideConvo() {
	Provide(svc_convo.NewConvoService)
}
