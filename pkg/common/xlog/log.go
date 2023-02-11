package xlog

import (
	"fmt"
	"log"
	"os"
)

func InitSystemLog(path string, app string) {
	logFile, err := os.OpenFile(path+app+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lmicroseconds | log.Ldate)
}
