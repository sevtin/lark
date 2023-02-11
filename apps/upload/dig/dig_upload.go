package dig

import (
	"lark/apps/upload/internal/service"
)

func provideUpload() {
	container.Provide(service.NewUploadService)
}
