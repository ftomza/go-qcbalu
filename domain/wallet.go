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

type DeliveryRPCWallet string

const (
	DeliveryRPCWalletAddBalanceByUserId    DeliveryRPCWallet = "add_balance_by_user_id"
	DeliveryRPCWalletDelBalanceByUserId    DeliveryRPCWallet = "del_balance_by_user_id"
	DeliveryRPCWalletGetBalanceByUserId    DeliveryRPCWallet = "get_balance_by_user_id"
	DeliveryRPCWalletDebtBalanceByUserId   DeliveryRPCWallet = "debt_balance_by_user_id"
	DeliveryRPCWalletCredBalanceByUserId   DeliveryRPCWallet = "cred_balance_by_user_id"
	DeliveryRPCWalletLockBalanceByUserId   DeliveryRPCWallet = "lock_balance_by_user_id"
	DeliveryRPCWalletUnlockBalanceByUserId DeliveryRPCWallet = "unlock_balance_by_user_id"
)

func (e DeliveryRPCWallet) IsValid() bool {
	switch e {
	case DeliveryRPCWalletAddBalanceByUserId,
		DeliveryRPCWalletDelBalanceByUserId,
		DeliveryRPCWalletGetBalanceByUserId,
		DeliveryRPCWalletDebtBalanceByUserId,
		DeliveryRPCWalletCredBalanceByUserId,
		DeliveryRPCWalletLockBalanceByUserId,
		DeliveryRPCWalletUnlockBalanceByUserId:
		return true
	}
	return false
}

func (e DeliveryRPCWallet) String() string {
	return string(e)
}

type EventNameWallet string

const (
	EventNameWalletBalanceSelf   EventNameWallet = "self"
	EventNameWalletBalanceDebt   EventNameWallet = "debt"
	EventNameWalletBalanceCred   EventNameWallet = "cred"
	EventNameWalletBalanceLock   EventNameWallet = "lock"
	EventNameWalletBalanceUnlock EventNameWallet = "unlock"
)

func (e EventNameWallet) IsValid() bool {
	switch e {
	case EventNameWalletBalanceDebt,
		EventNameWalletBalanceCred,
		EventNameWalletBalanceLock,
		EventNameWalletBalanceUnlock:
		return true
	}
	return false
}

func (e EventNameWallet) String() string {
	return string(e)
}

type EventWallet string

const (
	EventWalletBalanceAdd EventWallet = "balance.add"
	EventWalletBalanceDel EventWallet = "balance.del"
	EventWalletBalanceUpd EventWallet = "balance.upd"
)

func (e EventWallet) IsValid() bool {
	switch e {
	case EventWalletBalanceAdd,
		EventWalletBalanceDel,
		EventWalletBalanceUpd:
		return true
	}
	return false
}

func (e EventWallet) String() string {
	return string(e)
}

type Wallet struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Balance    int       `json:"balance"`
	Lock       bool      `json:"lock"`
	CreateTime time.Time `json:"create_at"`
	UpdateTime time.Time `json:"update_at"`
	Version    string    `json:"version"`
}

type WalletUserIDRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type WalletBalanceRequest struct {
	WalletUserIDRequest
	Sum int `json:"sum"`
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
