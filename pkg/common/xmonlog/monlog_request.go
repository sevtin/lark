package xmonlog

import (
	"time"
)

type RequestLog struct {
	CreatedAt   time.Time `bson:"created_at"`
	UserID      string    `bson:"user_id"`
	ClientIP    string    `bson:"client_ip"`
	Method      string    `bson:"method"`
	Uri         string    `bson:"uri"`
	Code        int       `bson:"code"`
	LatencyTime string    `bson:"latency_time"`
	UserAgent   string    `bson:"user_agent"`
	Params      string    `bson:"params"`
}

func SaveRequestLog(log RequestLog) {
	Insert("request_log", log)
}

type RequestLogErr struct {
	CreatedAt time.Time `bson:"created_at"`
	UserID    string    `bson:"user_id"`
	ClientIP  string    `bson:"client_ip"`
	Method    string    `bson:"method"`
	Uri       string    `bson:"uri"`
	Code      int       `bson:"code"`
	Err       string    `bson:"err"`
}

func SaveReqErrLog(log RequestLogErr) {
	Insert("request_err_log", log)
}
