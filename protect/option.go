package protect

import (
	"context"
	"time"
)

type Options struct {
	//base64 encoded
	RemoteURL string
	//base64 encoded
	RemoteSecret string
	//base64 encoded
	RemoteKey string
	Backoff   func() time.Duration
	stopper   func(ctx context.Context) error
	RemoteProvider
}
