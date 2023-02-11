package dig

import (
	"go.uber.org/dig"
	"lark/apps/cloud_msg/internal/config"
	"lark/apps/cloud_msg/internal/server"
	"lark/apps/cloud_msg/internal/server/cloud_msg"
	"lark/apps/cloud_msg/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(cloud_msg.NewCloudMessageServer)
	container.Provide(service.NewCloudMessageService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
