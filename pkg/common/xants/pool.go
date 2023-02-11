package xants

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Pool struct {
	wg *sync.WaitGroup
	pf *ants.PoolWithFunc
}

func NewPoolWithFunc(size int, f func(interface{})) (pool *Pool, err error) {
	var (
		pf *ants.PoolWithFunc
	)
	pool = &Pool{}
	pf, err = ants.NewPoolWithFunc(size, func(args interface{}) {
		f(args)
		pool.wg.Done()
	})
	if err != nil {
		return
	}
	pool.wg = &sync.WaitGroup{}
	pool.pf = pf
	return
}

func (p *Pool) Invoke(args interface{}) (err error) {
	p.wg.Add(1)
	err = p.pf.Invoke(args)
	return
}

func (p *Pool) Wait() {
	p.wg.Wait()
	p.pf.Release()
}
