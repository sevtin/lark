package dig

import (
	"lark/apps/interfaces/internal/service/svc_lbs"
)

func provideLbs() {
	container.Provide(svc_lbs.NewLbsService)
}
