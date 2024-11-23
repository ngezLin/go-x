package messaging

import "context"

type Subscriber interface {
	Subscribe(ctx context.Context, subs ...Topic) (err error)
}

type SubscriberCloser interface {
	CloseSubscriber() error
}
