package xcron

import (
	"log/slog"
	"testing"
	"time"
)

func TestTaskJob(t *testing.T) {
	args := &TaskArgs{
		Key:  "sync_message",
		Name: "同步消息",
		Spec: "CRON_TZ=Asia/Shanghai 30 08 ? * * *",
	}
	args = &TaskArgs{
		Key:  "update_message",
		Name: "更新消息",
		Spec: "*/5 * * * * *",
	}
	job := TaskJob{Params: nil, Task: args, Func: func(args interface{}) (err error) {
		slog.Info(time.Now().Format("2006-01-02 15:04:05"))
		return
	}}
	_, err := Scheduler.AddJob(job.Task, job)
	if err != nil {
		slog.Warn(err.Error())
	}
	time.Sleep(time.Hour)
}

func TestTaskFunc(t *testing.T) {
	args := &TaskArgs{
		Key:  "update_message",
		Name: "更新消息",
		Spec: "*/5 * * * * *",
	}
	_, err := Scheduler.AddFunc(args, func() {
		slog.Info(time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		slog.Warn(err.Error())
	}
	time.Sleep(time.Hour)
}
