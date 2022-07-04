package pool

import (
	"context"

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
	}
	for i := 0; i < concurrenceCount; i++ {
		p.ch <- struct{}{}
	}
	return p
}

type pool struct {
	ch               chan struct{}
	concurrenceCount int
}

func (p *pool) Go(f func()) {
	p.CtxGo(context.TODO(), f)
}

func (p *pool) CtxGo(ctx context.Context, f func()) {
	<-p.ch
	go func() {
		dev.Recover(recover())
		f()
		p.ch <- struct{}{}
	}()
}
