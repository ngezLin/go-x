package messaging

import "context"

// SyncPublisherWithKeyAndHeader is interface for handle sync publisher with key and header.
type SyncPublisherWithKeyAndHeader interface {

	// PublishSyncWithKeyAndHeader will publish the body in sync with key and header.
	// The default header will be attached:
	// 1. content-type -> default: application/json
	// 2. schema-version -> default: 1.0
	// 3. Traceparent -> only if you integrate well with new relic.
	// 4. origin -> if you specified when create new instance of the implementation.
	PublishSyncWithKeyAndHeader(ctx context.Context, topic string, key string, body interface{}, headers map[string]interface{}, opts ...PublishOption) error

	// PublishSyncWithKey will publish the body in sync with key and without header.
	// The default header will be attached:
	// 1. content-type -> default: application/json
	// 2. schema-version -> default: 1.0
	// 3. Traceparent -> only if you integrate well with new relic.
	// 4. origin -> if you specified when create new instance of the implementation.
	PublishSyncWithKey(ctx context.Context, topic string, key string, body interface{}, opts ...PublishOption) error
}

// AsyncPublisherWithKeyAndHeader is interface for handle async publisher with key and header.
type AsyncPublisherWithKeyAndHeader interface {

	// PublishAsyncWithKeyAndHeader will publish the body in sync with key and header.
	// This function will return *Promise,
	// you can handle the callback on success or on error by calling the function Then(callback func(err error)).
	// If error, the err will not nil. If success, the err will be nil.
	// The default header will be attached:
	// 1. content-type -> default: application/json
	// 2. schema-version -> default: 1.0
	// 3. Traceparent -> only if you integrate well with new relic.
	// 4. origin -> if you specified when create new instance of the implementation.
	PublishAsyncWithKeyAndHeader(ctx context.Context, topic string, key string, body interface{}, headers map[string]interface{}, opts ...PublishOption) (*Promise, error)

	// PublishAsyncWithKey will publish the body in sync with key and without header.
	// This function will return *Promise,
	// you can handle the callback on success or on error by calling the function Then(callback func(err error)).
	// If error, the err will not nil. If success, the err will be nil.
	// The default header will be attached:
	// 1. content-type -> default: application/json
	// 2. schema-version -> default: 1.0
	// 3. Traceparent -> only if you integrate well with new relic.
	// 4. origin -> if you specified when create new instance of the implementation.
	PublishAsyncWithKey(ctx context.Context, topic string, key string, body interface{}, opts ...PublishOption) (*Promise, error)
}

// PublisherWithKeyAndHeader is interface that contains contract for sync and async publisher with key and header.
type PublisherWithKeyAndHeader interface {
	AsyncPublisherWithKeyAndHeader
	SyncPublisherWithKeyAndHeader
	PublisherCloser
}

// PublisherTransaction is interface to handle publisher transactional API.
type PublisherTransaction interface {
	IsTransactional() bool
	BeginTxn() error
	CommitTxn() error
	AbortTxn() error
	StatusTxn() PublisherTxnStatusFlag
	AddMessageToTxn(msg Message) error
}

// PublisherCloser is interface for publisher close function.
type PublisherCloser interface {
	// ClosePublisher will gracefully close the publisher and wait until all the messages are published.
	ClosePublisher() error
}
