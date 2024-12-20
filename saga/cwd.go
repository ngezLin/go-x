package saga

import (
	"context"
	"time"
)

type cwd struct {
	ctx context.Context
}

func (*cwd) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*cwd) Done() <-chan struct{}       { return nil }
func (*cwd) Err() error                  { return nil }

func (l *cwd) Value(key interface{}) interface{} {
	return l.ctx.Value(key)
}

func NewCopyContext(ctx context.Context) *cwd {
	return &cwd{ctx}
}
