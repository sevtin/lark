package main

import (
	"lark/pkg/conf"
	"lark/scripts/generate-gorm/gengorm"
)

func main() {
	cfg := &conf.Mysql{
		Address:      "127.0.0.1",
		Username:     "root",
		Password:     "",
		Db:           "canary",
		MaxOpenConns: 20,
		MaxIdleConns: 10,
		MaxLifetime:  120000,
		MaxIdleTime:  0,
		Charset:      "utf8mb4",
		Debug:        false,
	}
	gengorm.GenGorm(cfg, "users", "./domain/po/")
}
