package xtimer

import (
	"fmt"
	"time"
)

type Timer struct {
	id       int64
	start    time.Time
	end      time.Time
	duration time.Duration
}

func NewTimer(d time.Duration) *Timer {
	t := &Timer{duration: d}
	return t
}

func (t *Timer) Run() {
	t.start = time.Now()
	t.id = t.start.Unix()
	t.end = t.start
	go func() {
		ticker := time.NewTicker(t.duration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println(t.id, " 时差(毫秒):", t.end.Sub(t.start).Milliseconds())
			}
		}
	}()
}

func (t *Timer) UpdateEndTime() {
	t.end = time.Now()
}
