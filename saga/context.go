package saga

import (
	"context"
	"errors"
	"sync"
)

type sagaCtxKey struct{}

var (
	SagaKey = sagaCtxKey{}
)

func collectSaga(ctx context.Context) (*saga, error) {
	deps, ok := ctx.Value(SagaKey).(*saga)
	if !ok {
		return nil, errors.New(ErrorSagaNotFound)
	}

	return deps, nil
}
func setSaga(ctx context.Context, deps *saga) context.Context {
	return context.WithValue(ctx, SagaKey, deps)
}

func BeginContext(ctx context.Context) (*saga, context.Context) {
	var (
		deps = &saga{
			wg:       &sync.WaitGroup{},
			chRevoke: make(chan func(context.Context) error),
		}
	)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			close(deps.chRevoke)
		case revoke := <-deps.chRevoke:
			deps.wg.Add(1)
			//running revokers
			go func() {
				err := revoke(ctx)
				if err != nil {
					if deps.errHandler != nil {
						deps.errHandler(ctx, err)
					}
				}
			}()
		}
	}(ctx)

	ctx = context.WithValue(ctx, SagaKey, deps)
	return deps, ctx
}

func RegisterContext(ctx context.Context, f func(context.Context) error) (context.Context, error) {
	deps, err := collectSaga(ctx)
	if err != nil {
		return ctx, err
	}
	deps.Register(ctx, f)
	ctx = setSaga(ctx, deps)
	return ctx, err
}

func SetErrorHandlerContext(ctx context.Context, f func(context.Context, error)) (context.Context, error) {
	deps, err := collectSaga(ctx)
	if err != nil {
		return ctx, err
	}
	deps.SetErrorHandler(ctx, f)
	ctx = setSaga(ctx, deps)
	return ctx, err
}
