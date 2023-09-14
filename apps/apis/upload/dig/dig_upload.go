package dig

import (
	"lark/apps/apis/upload/internal/service"
)

func provideUpload() {
	Provide(service.NewUploadService)
}
