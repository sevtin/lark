package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string              `yaml:"name"`
	ServerID         int                 `yaml:"server_id"`
	Log              string              `yaml:"log"`
	Monitor          *conf.Monitor       `yaml:"monitor"`
	GrpcServer       *conf.Grpc          `yaml:"grpc_server"`
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Mysql            *conf.Mysql         `yaml:"mysql"`
	Mongo            *conf.Mongo         `yaml:"mongo"`
	Redis            *conf.Redis         `yaml:"redis"`
	MsgGatewayServer *conf.GrpcServer    `yaml:"msg_gateway_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	Platforms        []*conf.Platform    `yaml:"platforms"`
	MsgConsumer      *conf.KafkaConsumer `yaml:"msg_consumer"`
	Elasticsearch    *conf.Elasticsearch `yaml:"elasticsearch"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/test/test.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	grpcPort = flag.Int("gp", 19999, "grpc server port")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.GrpcServer.ServerID = config.ServerID
	config.GrpcServer.Port = *grpcPort

	xlog.Shared(config.Log, config.Name+utils.IntToStr(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
