package saga

import (
	"context"
	"sync"
)

type (
	saga struct {
		wg         *sync.WaitGroup
		chRevoke   chan func(context.Context) error
		errHandler func(context.Context, error)
		queues     int
		mux        *sync.Mutex
	}
	Saga interface {
		Done(ctx context.Context)
		ErrorHandler(ctx context.Context, handler func(context.Context, error)) *saga
	}
)

func Begin(ctx context.Context) *saga {
	var (
		deps = &saga{
			wg:       &sync.WaitGroup{},
			chRevoke: make(chan func(context.Context) error),
			mux:      &sync.Mutex{},
			queues:   0,
		}
	)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			close(deps.chRevoke)
		case revoke := <-deps.chRevoke:
			//running revokers
			go func() {
				deps.wg.Wait()
				err := revoke(ctx)
				if err != nil {
					if deps.errHandler != nil {
						deps.errHandler(ctx, err)
					}
				}
			}()
		}
	}(ctx)

	return deps
}

func (dep *saga) Done() {
	for dep.queues > 0 {
		dep.wg.Done()
		dep.queues--
	}
	close(dep.chRevoke)
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
