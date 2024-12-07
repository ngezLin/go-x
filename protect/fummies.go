package protect

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type fummies struct {
	key   string
	child *fummies
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
