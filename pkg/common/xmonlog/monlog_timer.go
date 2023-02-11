package xmonlog

import (
	"time"
)

type TimerLog struct {
	CreatedAt   time.Time   `bson:"created_at"`
	Key         string      `bson:"key"`
	Name        string      `bson:"name"`
	Spec        string      `bson:"spec"`
	LatencyTime string      `bson:"latency_time"`
	Params      interface{} `bson:"params"`
	Err         string      `bson:"err"`
}

func SaveTimerLog(log TimerLog) {
	Insert("timer_log", log)
}
