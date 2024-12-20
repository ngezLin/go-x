package saga

import (
	"context"
	"sync"
)

type (
	saga struct {
		wg          *sync.WaitGroup
		chRevoke    chan func(context.Context) error
		chCompleted chan int
		errHandler  func(context.Context, error)
		queues      int
		mux         *sync.Mutex
	}
	Saga interface {
		Done(ctx context.Context)
		SetErrorHandler(ctx context.Context, handler func(context.Context, error)) *saga
		Register(ctx context.Context, f func(context.Context) error)
	}
)

func Begin(ctx context.Context) *saga {
	var (
		deps = &saga{
			wg:          &sync.WaitGroup{},
			chRevoke:    make(chan func(context.Context) error),
			chCompleted: make(chan int),
			mux:         &sync.Mutex{},
			queues:      0,
		}
	)

	go func() {
		for {
			select {
			case revoke, ok := <-deps.chRevoke:
				if !ok {
					return
				}
				go func() {
					cCtx := NewCopyContext(ctx)
					deps.wg.Wait()
					err := revoke(cCtx)
					if err != nil && deps.errHandler != nil {
						deps.errHandler(cCtx, err)
					}
				}()
			case <-ctx.Done():
				close(deps.chRevoke)
				return
			}
		}
	}()

	return deps
}

func (dep *saga) Done() {
	close(dep.chRevoke)
	dep.mux.Lock()
	defer dep.mux.Unlock()
	for dep.queues > 0 {
		dep.wg.Done()
		dep.queues--
	}
}

func (dep *saga) SetErrorHandler(ctx context.Context, f func(context.Context, error)) *saga {
	dep.errHandler = f
	return dep
}

func (dep *saga) Register(ctx context.Context, f func(context.Context) error) {
	dep.wg.Add(1)
	dep.mux.Lock()
	defer dep.mux.Unlock()

	dep.queues++
	dep.chRevoke <- f
}
