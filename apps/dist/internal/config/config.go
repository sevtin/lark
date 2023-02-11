package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string              `yaml:"name"`
	ServerID         int                 `yaml:"server_id"`
	Log              string              `yaml:"log"`
	Monitor          conf.Monitor        `yaml:"monitor"`
	GrpcServer       *conf.Grpc          `yaml:"grpc_server"`
	MsgGatewayServer *conf.GrpcServer    `yaml:"msg_gateway_server"`
	ChatMemberServer *conf.GrpcServer    `yaml:"chat_member_server"`
	Etcd             *conf.Etcd          `yaml:"etcd"`
	Platforms        []*conf.Platform    `yaml:"platforms"`
	GwMsgProducer    *conf.KafkaProducer `yaml:"gw_msg_producer"`
	MsgConsumer      *conf.KafkaConsumer `yaml:"msg_consumer"`
	Redis            *conf.Redis         `yaml:"redis"`
}

var (
	config = new(Config)
)

var (
	confFile    = flag.String("cfg", "./configs/distribute.yaml", "config file")
	serverId    = flag.Int("sid", 1, "server id")
	grpcPort    = flag.Int("gp", 7400, "grpc server port")
	monitorPort = flag.Int("mp", 7401, "metrics server port")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.GrpcServer.ServerID = config.ServerID
	config.GrpcServer.Port = *grpcPort
	config.Monitor.Port = *monitorPort

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+utils.IntToStr(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
