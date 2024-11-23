package messaging

import "context"

type Message interface {
	Bind(input any) error
	Context() context.Context
	WithContext(ctx context.Context) Message
	ContextCancel()
	GetMessageClaim() any
}
