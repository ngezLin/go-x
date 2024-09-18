package httproundtripper

import (
	"net/http"
)

type (
	loggingRoundTripper struct {
		next http.RoundTripper
		opt  *Option
	}
	LoggingRoundTripper interface {
		http.RoundTripper
	}
)

func NewLogging(next http.RoundTripper, opt *Option) *loggingRoundTripper {
	return &loggingRoundTripper{
		next: next,
	}
}

func (s *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return s.next.RoundTrip(req)
}
