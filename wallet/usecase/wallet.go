package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	iContext "github.com/ftomza/go-qcbalu/pkg/context"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/google/uuid"
)

type options struct {
	timeout time.Duration
}

type Option func(*options)

func SetTimeout(duration time.Duration) Option {
	return func(o *options) {
		o.timeout = duration
	}
}

type walletUsecase struct {
	walletRepository domain.WalletRepository
	options
}

func NewWalletUsecase(w domain.WalletRepository, opts ...Option) (domain.WalletUsecase, error) {
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}

	if w == nil {
		return nil, errors.New("ucase/wallet: repository not set")
	}

	return &walletUsecase{
		walletRepository: w,
		options:          options,
	}, nil
}

func (w *walletUsecase) AddByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {

		item = &domain.Wallet{
			ID:      uuid.New(),
			UserID:  userID,
			Balance: sum,
		}
		item, err = w.walletRepository.Store(ctx, item)
		if err != nil {
			return fmt.Errorf("ucase/wallet: %w", err)
		}
		return nil
	})
	return item, err
}

func (w *walletUsecase) DelByUserID(ctx context.Context, userID uuid.UUID) (err error) {
	return w.execWithTimeout(ctx, func(ctx context.Context) error {

		err = w.walletRepository.DeleteByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("ucase/wallet: %w", err)
		}
		return nil
	})
}

func (w *walletUsecase) GetByUserID(ctx context.Context, userID uuid.UUID) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {

		item, err = w.walletRepository.GetByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("ucase/wallet: %w", err)
		}
		return nil
	})
	return item, err
}

func (w *walletUsecase) DebtBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {
		return ExecTrans(ctx, w.walletRepository.GetTransaction(ctx), func(ctx context.Context) error {
			var itemStore *domain.Wallet
			itemStore, err = w.getItemByUserIDAndCheckLock(ctx, userID)
			if err != nil {
				return err
			}

			itemStore.Balance = itemStore.Balance - sum
			if itemStore.Balance < 0 {
				return domain.ErrInsufficientFunds
			}

			item, err = w.walletRepository.Update(ctx, item)
			return err
		})
	})
	return item, err
}

func (w *walletUsecase) CredBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {
		return ExecTrans(ctx, w.walletRepository.GetTransaction(ctx), func(ctx context.Context) error {
			var itemStore *domain.Wallet
			itemStore, err = w.getItemByUserIDAndCheckLock(ctx, userID)
			if err != nil {
				return err
			}

			itemStore.Balance = itemStore.Balance + sum

			item, err = w.walletRepository.Update(ctx, item)
			return err
		})
	})
	return item, err
}

func (w *walletUsecase) LockByUserID(ctx context.Context, userID uuid.UUID) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {
		return ExecTrans(ctx, w.walletRepository.GetTransaction(ctx), func(ctx context.Context) error {

			item, err = w.walletRepository.GetByUserID(ctx, userID)
			if err != nil {
				return fmt.Errorf("ucase/wallet: %w", err)
			}

			if item.Lock {
				return nil
			}

			item.Lock = true

			item, err = w.walletRepository.Update(ctx, item)
			return err
		})
	})
	return item, err
}

func (w *walletUsecase) UnlockByUserID(ctx context.Context, userID uuid.UUID) (item *domain.Wallet, err error) {
	err = w.execWithTimeout(ctx, func(ctx context.Context) error {
		return ExecTrans(ctx, w.walletRepository.GetTransaction(ctx), func(ctx context.Context) error {

			item, err = w.walletRepository.GetByUserID(ctx, userID)
			if err != nil {
				return fmt.Errorf("ucase/wallet: %w", err)
			}

			if !item.Lock {
				return nil
			}

			item.Lock = false

			item, err = w.walletRepository.Update(ctx, item)
			return err
		})
	})
	return item, err
}

func (w *walletUsecase) execWithTimeout(ctx context.Context, fn func(ctx context.Context) error) error {
	return iContext.ExecWithTimeout(ctx, w.timeout, fn)
}

func (w *walletUsecase) getItemByUserIDAndCheckLock(ctx context.Context, userID uuid.UUID) (item *domain.Wallet, err error) {
	item, err = w.walletRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("ucase/wallet: %w", err)
	}
	if item.Lock {
		return nil, domain.ErrLocked
	}
	return
}
