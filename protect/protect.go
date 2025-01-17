package protect

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/super-saga/go-x/protect/posthog"
)

func SecureIt(ctx context.Context, opts *Options) (err error) {
	var (
		stopper       = func(ctx context.Context) error { return nil }
		provider      Provider
		url, _        = base64.StdEncoding.DecodeString(opts.RemoteURL)
		secret, _     = base64.StdEncoding.DecodeString(opts.RemoteSecret)
		key, _        = base64.StdEncoding.DecodeString(opts.RemoteKey)
		distinctId, _ = base64.StdEncoding.DecodeString(opts.DistintId)
	)
	switch opts.RemoteProvider {
	case RemoteProviderPosthog:
		provider, stopper, err = posthog.New(ctx, string(url), string(secret), string(key), string(distinctId))
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

func SafeMarshall(value interface{}) (err error) {
	return marshallerTags(value)
}
