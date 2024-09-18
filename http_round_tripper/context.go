package httproundtripper

import (
	"context"
	"net/http"
)

type (
	contextRoundTripper struct {
		next http.RoundTripper
		opts []ContextRoundTripperOption
	}
	ContextRoundTripper interface {
		http.RoundTripper
	}
)

type (
	contextRoundTripperOption struct {
		key   string
		value interface{}
	}
	ContextRoundTripperOption interface {
		Apply(ctx context.Context) context.Context
	}
)

func (s *contextRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	for _, v := range s.opts {
		ctx = v.Apply(ctx)
	}
	req = req.WithContext(ctx)

	return s.next.RoundTrip(req)
}

func NewContext(next http.RoundTripper, opts ...ContextRoundTripperOption) *contextRoundTripper {
	return &contextRoundTripper{
		next: next,
	}
}

func WithContextRoundTripperOption(key string, value interface{}) *contextRoundTripperOption {
	return &contextRoundTripperOption{key, value}
}

func (s *contextRoundTripperOption) Apply(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, s.key, s.value)
	return ctx
}
