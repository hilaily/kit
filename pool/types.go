package pool

import "context"

type IPool interface {
	// SetCap sets the goroutine capacity of the pool.
	//SetCap(cap int32)
	// Go executes f.
	Go(f func())
	// CtxGo executes f and accepts the context.
	CtxGo(ctx context.Context, f func())
	// SetPanicHandler sets the panic handler.
	//SetPanicHandler(f func(context.Context, interface{}))
}
