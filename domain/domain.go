package domain

import "context"

type Transaction interface {
	Begin(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, trans interface{}) error
	Rollback(ctx context.Context, trans interface{}) error
}

type EventHeaders struct {
	Event string
}
