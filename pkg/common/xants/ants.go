package xants

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Ants struct {
	wg *sync.WaitGroup
	pf func()
}

func NewAnts(pf func()) (s *Ants) {
	return &Ants{
		wg: &sync.WaitGroup{},
		pf: pf,
	}
}

func (s *Ants) Submit() {
	s.wg.Add(1)
	_ = ants.Submit(s.pf)
}

func (s *Ants) Wait() {
	s.wg.Wait()
	ants.Release()
}

func Submit(task func()) (err error) {
	err = ants.Submit(task)
	return
}
