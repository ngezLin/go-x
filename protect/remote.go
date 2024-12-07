package protect

import "context"

type RemoteProvider string

var RemoteProviderPosthog RemoteProvider = "posthog"

type Provider interface {
	Initialize(ctx context.Context, url, secret string) error
	Seek(ctx context.Context) (value bool)
}
