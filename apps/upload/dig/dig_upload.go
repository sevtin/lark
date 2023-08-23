package dig

import (
	"lark/apps/upload/internal/service"
)

func provideUpload() {
	Provide(service.NewUploadService)
}
