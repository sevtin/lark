package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string           `yaml:"name"`
	ServerID         int              `yaml:"server_id"`
	Log              string           `yaml:"log"`
	Etcd             *conf.Etcd       `yaml:"etcd"`
	Redis            *conf.Redis      `yaml:"redis"`
	MsgGatewayServer *conf.GrpcServer `yaml:"msg_gateway_server"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/server_mgr.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+utils.IntToStr(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
