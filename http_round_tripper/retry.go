package http_round_tripper

import (
	"net/http"
	"time"
)

var (
	DefaultRetriesOptions = retryRoundTripperOption{
		isRetryable: IsInternalError,
		maxRetries:  5,
		backOff: func() time.Duration {
			return time.Second * 2
		},
		after: time.After,
	}
	DefaultMaxRetries = 5
)

type (
	retryRoundTripper struct {
		next http.RoundTripper
		opts *retryRoundTripperOption
	}
	RetryRoundTripper interface {
		http.RoundTripper
	}
)

type (
	retryRoundTripperOption struct {
		maxRetries  int
		backOff     func() time.Duration
		isRetryable func(res *http.Response) bool
		after       func(d time.Duration) <-chan time.Time
	}
	retryRoundTripperOptions func(*retryRoundTripperOption)
)

func IsInternalError(res *http.Response) bool {
	return res.StatusCode == http.StatusInternalServerError
}

func (s *retryRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	var attempts int
	for {
		res, err := s.next.RoundTrip(req)
		attempts++

		if attempts == s.opts.maxRetries {
			return res, err
		}

		if !s.opts.isRetryable(res) {
			return res, err
		}

		select {
		case <-req.Context().Done():
			return res, req.Context().Err()
		case <-time.After(s.opts.backOff()):
		}
	}
}

func (s *retryRoundTripper) WithBackoff(f func() time.Duration) *retryRoundTripperOption {
	s.opts.backOff = f
	return s.opts
}

func (s *retryRoundTripper) WithIsRetryable(f func(res *http.Response) bool) *retryRoundTripperOption {
	s.opts.isRetryable = f
	return s.opts
}

func (s *retryRoundTripper) WithMaxRetries(max int) *retryRoundTripperOption {
	s.opts.maxRetries = max
	return s.opts
}

func NewRetryRoundTripper(next http.RoundTripper, opts ...retryRoundTripperOptions) *retryRoundTripper {
	return &retryRoundTripper{
		next: next,
	}
}
