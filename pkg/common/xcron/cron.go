package xcron

import (
	"errors"
	"github.com/robfig/cron/v3"
	"log/slog"
	"sync"
)

var (
	ERR_TASK_ARGS_INVALID = errors.New("task args invalid")
)

var (
	Scheduler Timer
)

type Timer interface {
	AddFunc(args *TaskArgs, function func()) (entryID cron.EntryID, err error)
	AddJob(args *TaskArgs, job interface{ Run() }) (entryID cron.EntryID, err error)
	GetTask(key string) (*Task, bool)
	StartTask(key string)
	StopTask(key string)
	Remove(key string, entryID int)
	Clear(key string)
	Close()
}

// 默认Job接口
type TaskJob struct {
	Params interface{}
	Task   *TaskArgs
	Func   func(args interface{}) (err error)
}

func (job TaskJob) Run() {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("task job run error:", err)
		}
	}()
	job.Func(job.Params)
}

type TaskArgs struct {
	Key  string
	Name string
	Spec string
}

// 小写开头字段无须外部赋值
type Task struct {
	*TaskArgs
	isJob bool
	cron  *cron.Cron
}

// timer 定时任务管理
type timer struct {
	tasks map[string]*Task
	sync.Mutex
}

func init() {
	Scheduler = newTimer()
}

func newTimer() Timer {
	return &timer{tasks: make(map[string]*Task)}
}

// 通过函数的方法添加任务
func (t *timer) AddFunc(args *TaskArgs, function func()) (entryID cron.EntryID, err error) {
	return t.addTask(args, false, function)
}

// 通过接口的方法添加任务
func (t *timer) AddJob(args *TaskArgs, job interface{ Run() }) (entryID cron.EntryID, err error) {
	return t.addTask(args, true, job)
}

// 校验任务参数
func (t *timer) validateArgs(args *TaskArgs, cmd interface{}) (err error) {
	if args.Key == "" || args.Spec == "" || cmd == nil {
		err = ERR_TASK_ARGS_INVALID
		return
	}
	return
}

func (t *timer) addTask(args *TaskArgs, isJob bool, cmd interface{}) (entryID cron.EntryID, err error) {
	if err = t.validateArgs(args, cmd); err != nil {
		return
	}
	t.Lock()
	defer t.Unlock()
	var (
		task *Task
		ok   bool
	)
	if task, ok = t.tasks[args.Key]; ok == false {
		task = &Task{
			TaskArgs: args,
			isJob:    isJob,
			cron:     cron.New(cron.WithSeconds()),
		}
		t.tasks[task.Key] = task
	}
	if isJob == true {
		entryID, err = task.cron.AddJob(args.Spec, cmd.(interface{ Run() }))
	} else {
		entryID, err = task.cron.AddFunc(args.Spec, cmd.(func()))
	}
	if err != nil {
		return
	}
	task.cron.Start()
	return
}

// GetTask 获取对应key的task 可能会为空
func (t *timer) GetTask(key string) (*Task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.tasks[key]
	return v, ok
}

// StartTask 开始任务
func (t *timer) StartTask(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.tasks[key]; ok {
		v.cron.Start()
	}
}

// StopTask 停止任务
func (t *timer) StopTask(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.tasks[key]; ok {
		v.cron.Stop()
	}
}

// Remove 删除指定任务
func (t *timer) Remove(key string, entryID int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.tasks[key]; ok {
		v.cron.Remove(cron.EntryID(entryID))
	}
}

// Clear 清除任务
func (t *timer) Clear(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.tasks[key]; ok {
		v.cron.Stop()
		delete(t.tasks, key)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.tasks {
		v.cron.Stop()
	}
}

/*
cron.Remove删除单个任务,cron.Stop停止整个cron调度。
cron.Remove通过EntryID指定任务,cron.Stop作用于全局。
cron.Remove只停止指定任务,cron.Stop会停止所有任务。
cron.Remove删除后需要再Add才能重新运行该任务,cron.Stop可以直接Start重新开始。
*/
