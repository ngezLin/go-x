package ctxdata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type Set func(ctx context.Context) context.Context
type SetMD func(md metadata.MD) metadata.MD

func Sets(ctx context.Context, fn ...Set) context.Context {
	for _, f := range fn {
		ctx = f(ctx)
	}
	return ctx
}

func SetsMD(md metadata.MD, fn ...SetMD) metadata.MD {
	for _, f := range fn {
		md = f(md)
	}
	return md
}
