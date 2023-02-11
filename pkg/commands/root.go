package commands

import (
	"errors"
	"flag"
	_ "lark/pkg/common/xstat"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	GMainInst MainInstance
	GSignal   chan os.Signal
)

type MainInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

func Run(inst MainInstance) {
	flag.Parse()

	if inst == nil {
		panic(errors.New("inst is nil, exit"))
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := inst.Initialize()
	if err != nil {
		panic(err)
		return
	}
	GMainInst = inst

	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
