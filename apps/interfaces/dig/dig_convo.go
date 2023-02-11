package dig

import (
	"lark/apps/interfaces/internal/service/svc_convo"
)

func provideConvo() {
	container.Provide(svc_convo.NewConvoService)
}
