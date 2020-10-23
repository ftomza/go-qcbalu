package repository

import (
	"context"
	"errors"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"
	"github.com/google/uuid"
)

type entWalletRepository struct {
	Conn *ent.Client
}

func NewEntWalletRepository(conn *ent.Client) domain.WalletRepository {
	return &entWalletRepository{conn}
}

func (e *entWalletRepository) Store(ctx context.Context, item *domain.Wallet) (newItem *domain.Wallet, err error) {
	if item == nil {
		return nil, ErrItemNotSet
	}

	err = ExecTrans(ctx, e.Conn, func(ctx context.Context, tx *ent.Tx) (err error) {
		var res *ent.Wallet
		res, err = e.Conn.Wallet.Create().
			SetID(item.ID).
			SetUserID(item.UserID).
			SetBalance(item.Balance).
			SetLock(item.Lock).
			Save(ctx)
		if err != nil {
			return err
		}

		newItem = res.ToDomainWallet()

		return nil
	})
	return
}

func (e *entWalletRepository) Update(ctx context.Context, item *domain.Wallet) (newItem *domain.Wallet, err error) {
	if item == nil {
		return nil, ErrItemNotSet
	}

	err = ExecTrans(ctx, e.Conn, func(ctx context.Context, tx *ent.Tx) (err error) {
		err = tx.Wallet.UpdateOneID(item.ID).
			SetUserID(item.UserID).
			SetBalance(item.Balance).
			SetLock(item.Lock).
			SetVersion(item.Version).
			Exec(ctx)
		if err != nil {
			return err
		}
		newItem, err = e.GetByUserID(ctx, item.UserID)
		return
	})

	return
}

func (e *entWalletRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (item *domain.Wallet, err error) {
	err = ExecTrans(ctx, e.Conn, func(ctx context.Context, tx *ent.Tx) (err error) {
		query := tx.Wallet.Query().Where(wallet.UserID(userID))
		var queryItem *ent.Wallet
		if queryItem, err = query.Only(ctx); err == nil {
			item = queryItem.ToDomainWallet()
		}
		return err
	})
	return item, err
}

func (e *entWalletRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) (err error) {
	return ExecTrans(ctx, e.Conn, func(ctx context.Context, tx *ent.Tx) (err error) {
		_, err = e.Conn.Wallet.Delete().Where(wallet.UserID(userID)).Exec(ctx)

		if err != nil && errors.Is(err, &ent.NotFoundError{}) {
			return
		}
		return nil
	})
}

func (e *entWalletRepository) GetTransaction(_ context.Context) domain.Transaction {
	return NewEntTransaction(e.Conn)
}
