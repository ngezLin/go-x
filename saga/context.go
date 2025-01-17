package saga

import (
	"context"
	"errors"
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
	deps := Begin(ctx)
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

func DoneContext(ctx context.Context) error {
	deps, err := collectSaga(ctx)
	if err != nil {
		return err
	}
	go func() {
		defer deps.Stop()
		<-deps.done()
	}()
	return nil
}

func DoneContextSync(ctx context.Context) <-chan int {
	deps, _ := collectSaga(ctx)
	return deps.done()
}
