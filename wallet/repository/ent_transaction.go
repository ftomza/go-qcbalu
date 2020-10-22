package repository

import (
	"context"
	"errors"
	"fmt"

	iContext "github.com/ftomza/go-qcbalu/pkg/context"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent"
)

func NewEntTransaction(client *ent.Client) *EntTransaction {
	return &EntTransaction{client}
}

type EntTransaction struct {
	client *ent.Client
}

type TransactionFunc func(context.Context, *ent.Tx) error

func ExecTrans(ctx context.Context, c *ent.Client, fn TransactionFunc) (err error) {
	if trans, ok := iContext.FromTrans(ctx); ok {
		tx, ok := trans.(*ent.Tx)
		if ok {
			return fn(ctx, tx)
		}
	}

	entTrans := NewEntTransaction(c)
	trans, err := entTrans.Begin(ctx)
	if err != nil {
		return err
	}

	tx, _ := trans.(*ent.Tx)
	defer func() {
		if v := recover(); v != nil {
			_ = entTrans.Rollback(ctx, trans)
			panic(v)
		}
	}()
	err = fn(iContext.NewTrans(ctx, trans), tx)
	if err != nil {
		_ = entTrans.Rollback(ctx, trans)
		return err
	}
	return entTrans.Commit(ctx, trans)
}

func (a *EntTransaction) Begin(ctx context.Context) (trans interface{}, err error) {

	trans, err = a.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo/ent: Open transaction error: %w", err)
	}
	return
}

func (a *EntTransaction) Commit(_ context.Context, trans interface{}) (err error) {

	tx, ok := trans.(*ent.Tx)
	if !ok {
		return errors.New("repo/ent: Unknown transaction type")
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("repo/ent: Commit transaction error: %w", err)
	}
	return nil
}

func (a *EntTransaction) Rollback(_ context.Context, trans interface{}) (err error) {

	tx, ok := trans.(*ent.Tx)
	if !ok {
		return errors.New("repo/ent: Unknown transaction type")
	}

	err = tx.Rollback()
	if err != nil {
		return fmt.Errorf("repo/ent: Rollback transaction error: %w", err)
	}
	return nil
}
