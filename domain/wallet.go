package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInsufficientFunds = errors.New("wallet: Insufficient funds")
	ErrLocked            = errors.New("wallet: Item blocked")
)

type Wallet struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Balance    int       `json:"balance"`
	Lock       bool      `json:"lock"`
	CreateTime time.Time `json:"create_at"`
	UpdateTime time.Time `json:"update_at"`
	Version    string    `json:"version"`
}

type WalletUsecase interface {
	AddByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *Wallet, err error)
	DelByUserID(ctx context.Context, userID uuid.UUID) (err error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (item *Wallet, err error)
	DebtBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *Wallet, err error)
	CredBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (item *Wallet, err error)
	LockByUserID(ctx context.Context, userID uuid.UUID) (item *Wallet, err error)
	UnlockByUserID(ctx context.Context, userID uuid.UUID) (item *Wallet, err error)
}

type WalletRepository interface {
	Store(ctx context.Context, item *Wallet) (newItem *Wallet, err error)
	Update(ctx context.Context, item *Wallet) (newItem *Wallet, err error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (item *Wallet, err error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) (err error)
	GetTransaction(ctx context.Context) Transaction
}
