package usecase

import (
	"context"

	"github.com/ftomza/go-qcbalu/domain"
	iContext "github.com/ftomza/go-qcbalu/pkg/context"
)

type TransactionFunc func(context.Context) error

func ExecTrans(ctx context.Context, transaction domain.Transaction, fn TransactionFunc) error {
	if _, ok := iContext.FromTrans(ctx); ok {
		return fn(ctx)
	}
	transUCase, err := transaction.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(iContext.NewTrans(ctx, transUCase))
	if err != nil {
		_ = transaction.Rollback(ctx, transUCase)
		return err
	}
	return transaction.Commit(ctx, transUCase)
}
