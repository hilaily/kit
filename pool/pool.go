package pool

import (
	"context"
	"sync"

	"github.com/hilaily/kit/dev"
)

var (
	_ IPool = &pool{}
)

func NewPool(concurrenceCount int) IPool {
	if concurrenceCount <= 0 {
		concurrenceCount = 1
	}
	p := &pool{
		concurrenceCount: concurrenceCount,
		ch:               make(chan struct{}, concurrenceCount),
		wg:               &sync.WaitGroup{},
	}
	for i := 0; i < concurrenceCount; i++ {
		p.ch <- struct{}{}
	}
	return p
}

type pool struct {
	ch               chan struct{}
	wg               *sync.WaitGroup
	concurrenceCount int
}

func (p *pool) Go(f func()) {
	p.CtxGo(context.TODO(), f)
}

func (p *pool) CtxGo(ctx context.Context, f func()) {
	p.wg.Add(1)
	<-p.ch
	go func() {
		dev.Recover(recover())
		f()
		p.wg.Done()
		p.ch <- struct{}{}
	}()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
