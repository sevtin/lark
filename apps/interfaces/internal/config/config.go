package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name                string             `yaml:"name"`
	ServerID            int                `yaml:"server_id"`
	Port                int                `yaml:"port"`
	Log                 string             `yaml:"log"`
	Etcd                *conf.Etcd         `yaml:"etcd"`
	Redis               *conf.Redis        `yaml:"redis"`
	AuthServer          *conf.GrpcServer   `yaml:"auth_server"`
	UserServer          *conf.GrpcServer   `yaml:"user_server"`
	ChatMsgServer       *conf.GrpcServer   `yaml:"chat_msg_server"`
	MessageServer       *conf.GrpcServer   `yaml:"message_server"`
	LinkServer          *conf.GrpcServer   `yaml:"link_server"`
	ChatMemberServer    *conf.GrpcServer   `yaml:"chat_member_server"`
	ChatInviteServer    *conf.GrpcServer   `yaml:"chat_invite_server"`
	ChatServer          *conf.GrpcServer   `yaml:"chat_server"`
	AvatarServer        *conf.GrpcServer   `yaml:"avatar_server"`
	ConvoServer         *conf.GrpcServer   `yaml:"convo_server"`
	LbsServer           *conf.GrpcServer   `yaml:"lbs_server"`
	RedEnvServer        *conf.GrpcServer   `yaml:"red_env_server"`
	RedEnvReceiveServer *conf.GrpcServer   `yaml:"red_env_receive_server"`
	OrderServer         *conf.GrpcServer   `yaml:"order_server"`
	PaymentServer       *conf.GrpcServer   `yaml:"payment_server"`
	Jaeger              *conf.Jaeger       `yaml:"jaeger"`
	Google              *conf.GoogleOAuth2 `yaml:"google_oauth2"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/api_gateway.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	port     = flag.Int("p", 8088, "api gateway default listen port 8088")
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
