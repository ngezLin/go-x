package saga

import (
	"context"
	"sync"
)

var DEFAULT_MAX_GO = 5

type (
	saga struct {
		wg          *sync.WaitGroup
		chRevoke    chan func(context.Context) error
		chCompleted chan int
		errHandler  func(context.Context, error)
		queues      int
		mux         *sync.Mutex
		maxGo       int
	}
	Saga interface {
		Stop()
		Done() <-chan int
		SetErrorHandler(ctx context.Context, handler func(context.Context, error)) *saga
		Register(ctx context.Context, f func(context.Context) error)
	}
)

func (deps *saga) process(ctx context.Context) {
	var (
		wg   = sync.WaitGroup{}
		cCtx = NewCopyContext(ctx)
	)

	for i := 0; i < deps.maxGo; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case revoke, ok := <-deps.chRevoke:
					if !ok {
						return
					}
					go func() {
						defer wg.Done()
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
	}
	wg.Wait()
}

func Begin(ctx context.Context, opts ...Option) *saga {
	var (
		deps = &saga{
			wg:          &sync.WaitGroup{},
			chRevoke:    make(chan func(context.Context) error),
			chCompleted: make(chan int),
			mux:         &sync.Mutex{},
			queues:      0,
			maxGo:       DEFAULT_MAX_GO,
		}
	)

	for _, opt := range opts {
		opt(deps)
	}

	go func() {
		deps.process(ctx)
		deps.chCompleted <- 1
	}()

	return deps
}

func (dep *saga) Done() <-chan int {
	close(dep.chRevoke)
	dep.mux.Lock()
	defer dep.mux.Unlock()
	for dep.queues > 0 {
		dep.wg.Done()
		dep.queues--
	}
	return dep.chCompleted
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

func (dep *saga) Stop() {

	close(dep.chCompleted)
}
