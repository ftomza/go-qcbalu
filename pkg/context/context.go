package context

import (
	"context"
	"time"
)

type (
	transCtx struct{}
)

func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

func ExecWithTimeout(ctx context.Context, timeout time.Duration, fn func(ctx context.Context) error) error {
	if timeout > 0 {
		ctxNew, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return fn(ctxNew)
	}
	return fn(ctx)
}
