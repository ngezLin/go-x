package posthog

import (
	"context"
	"time"

	goposthog "github.com/posthog/posthog-go"
)

type posthog struct {
	client     goposthog.Client
	key        string
	distinctId string
}

func New(ctx context.Context, url, secret, key, distinctId string) (ref *posthog, stopper func(ctx context.Context) error, err error) {
	stopper = func(ctx context.Context) error { return nil }
	client, err := goposthog.NewWithConfig(secret, goposthog.Config{
		Endpoint: url,
		Interval: time.Minute * 30,
	})
	stopper = func(ctx context.Context) error { return client.Close() }
	ref = &posthog{
		client:     client,
		key:        key,
		distinctId: distinctId,
	}
	return
}

func (ref *posthog) Seek(ctx context.Context) bool {
	value, _ := ref.client.IsFeatureEnabled(
		goposthog.FeatureFlagPayload{
			Key:        ref.key,
			DistinctId: ref.distinctId,
		})
	return value.(bool)
}

func (ref *posthog) Initialize(ctx context.Context, url, secret string) (err error) {
	return
}
