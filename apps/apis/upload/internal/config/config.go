package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name       string           `yaml:"name"`
	ServerID   int              `yaml:"server_id"`
	Port       int              `yaml:"port"`
	Log        string           `yaml:"log"`
	Etcd       *conf.Etcd       `yaml:"etcd"`
	Redis      *conf.Redis      `yaml:"redis"`
	UserServer *conf.GrpcServer `yaml:"user_server"`
	ChatServer *conf.GrpcServer `yaml:"chat_server"`
	Minio      *conf.Minio      `yaml:"minio"`
	Jaeger     *conf.Jaeger     `yaml:"jaeger"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/upload.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	port     = flag.Int("p", 7800, "default listen port 7800")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.Port = *port

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+utils.IntToStr(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
