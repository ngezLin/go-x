package mail

import "context"

type SMTP interface {
	Prepare() *email
	Send(ctx context.Context, email *email) error
}
