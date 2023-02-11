package xmonlog

import (
	"time"
)

type MysqlLog struct {
	CreatedAt time.Time `bson:"created_at"`
	Src       string    `bson:"src"`
	Duration  float64   `bson:"duration"`
	Rows      int64     `bson:"rows"`
	Sql       string    `bson:"sql"`
}

type MysqlErrLog struct {
	CreatedAt time.Time `bson:"created_at"`
	Src       string    `bson:"src"`
	Duration  float64   `bson:"duration"`
	Sql       string    `bson:"sql"`
	Err       string    `bson:"err"`
}

func SaveMysqlLog(log MysqlLog) {
	Insert("mysql_log", log)
}

func SaveMysqlErrLog(log MysqlErrLog) {
	Insert("mysql_err_log", log)
}
