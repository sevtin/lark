package template

var InternalConfigCode = `
package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name             string             "yaml:"name""
	ServerID         int                "yaml:"server_id""
	Port             int                "yaml:"port""
	Log              string             "yaml:"log""
	Etcd             *conf.Etcd         "yaml:"etcd""
	Mysql            *conf.Mysql        "yaml:"mysql""
	Redis            *conf.Redis        "yaml:"redis""
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/api_{{.PackageName}}.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	port     = flag.Int("p", 6600, "default listen port 6600")
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
`
var InternalConfigTemplate = ParseTemplate(InternalConfigCode)
