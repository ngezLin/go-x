package saga

import (
	"context"
	"sync"
)

var DEFAULT_MAX_GO = 5

type (
	saga struct {
		runners     []func(context.Context) error
		chRunner    chan func(context.Context) error
		chCompleted chan int
		errHandler  func(context.Context, error)
		mux         *sync.Mutex
		maxGo       int
	}
	Saga interface {
		Done()
		DoneSync() <-chan int
		Stop()
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
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					close(deps.chRunner)
					return
				case run, ok := <-deps.chRunner:
					if !ok {
						return
					}
					err := run(cCtx)
					if err != nil && deps.errHandler != nil {
						deps.errHandler(cCtx, err)
					}
				}
			}
		}()
	}

	for _, runner := range deps.runners {
		deps.chRunner <- runner
	}
	close(deps.chRunner)

	wg.Wait()
}

func Begin(ctx context.Context, opts ...Option) *saga {
	var (
		deps = &saga{
			runners:     []func(context.Context) error{},
			chRunner:    make(chan func(context.Context) error),
			chCompleted: make(chan int),
			mux:         &sync.Mutex{},
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

func (dep *saga) Done() {
	go func() {
		defer dep.Stop()
		<-dep.done()
	}()
}

func (dep *saga) DoneSync() <-chan int {
	return dep.done()
}

func (dep *saga) done() <-chan int {
	close(dep.chRunner)
	return dep.chCompleted
}

func (dep *saga) SetErrorHandler(ctx context.Context, f func(context.Context, error)) *saga {
	dep.errHandler = f
	return dep
}

func (dep *saga) Register(ctx context.Context, f func(context.Context) error) {
	dep.mux.Lock()
	defer dep.mux.Unlock()
	dep.runners = append(dep.runners, f)
}

func (dep *saga) Stop() {
	close(dep.chCompleted)
}
