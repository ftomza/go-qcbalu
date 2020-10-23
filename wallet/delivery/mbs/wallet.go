package mbs

import (
	"context"

	"github.com/ftomza/go-qcbalu/pkg/mbs"

	"github.com/ftomza/go-qcbalu/domain"
)

type Wallet struct {
	walletUCase domain.WalletUsecase
	rpc         mbs.RPCService
	pub         mbs.PUBService
}

func NewMBSWallet(rpc mbs.RPCService, pub mbs.PUBService, wuc domain.WalletUsecase) *Wallet {
	rpcWallet := &Wallet{
		walletUCase: wuc,
		rpc:         rpc,
		pub:         pub,
	}

	rpcWallet.rpc.
		AddMethod(domain.DeliveryRPCWalletAddBalanceByUserId.String(), rpcWallet.addBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletDelBalanceByUserId.String(), rpcWallet.delBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletGetBalanceByUserId.String(), rpcWallet.getBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletDebtBalanceByUserId.String(), rpcWallet.debtBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletCredBalanceByUserId.String(), rpcWallet.credBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletLockBalanceByUserId.String(), rpcWallet.lockBalanceByUserId).
		AddMethod(domain.DeliveryRPCWalletUnlockBalanceByUserId.String(), rpcWallet.unlockBalanceByUserId)
	return rpcWallet
}

func (w Wallet) addBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletBalanceRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletBalanceRequest) (interface{}, error) {
		return w.walletUCase.AddByUserID(ctx, req.UserID, req.Sum)
	}, w.eventWallet(domain.EventNameWalletBalanceSelf, domain.EventWalletBalanceAdd))
}

func (w Wallet) delBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletUserIDRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletUserIDRequest) (interface{}, error) {
		return &domain.Wallet{UserID: req.UserID}, w.walletUCase.DelByUserID(ctx, req.UserID)
	}, w.eventWallet(domain.EventNameWalletBalanceSelf, domain.EventWalletBalanceDel))
}

func (w Wallet) getBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletUserIDRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletUserIDRequest) (interface{}, error) {
		return w.walletUCase.GetByUserID(ctx, req.UserID)
	})
}

func (w Wallet) debtBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletBalanceRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletBalanceRequest) (interface{}, error) {
		return w.walletUCase.DebtBalanceByUserID(ctx, req.UserID, req.Sum)
	}, w.eventWallet(domain.EventNameWalletBalanceDebt, domain.EventWalletBalanceUpd))
}

func (w Wallet) credBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletBalanceRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletBalanceRequest) (interface{}, error) {
		return w.walletUCase.CredBalanceByUserID(ctx, req.UserID, req.Sum)
	}, w.eventWallet(domain.EventNameWalletBalanceCred, domain.EventWalletBalanceUpd))
}

func (w Wallet) lockBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletUserIDRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletUserIDRequest) (interface{}, error) {
		return w.walletUCase.LockByUserID(ctx, req.UserID)
	}, w.eventWallet(domain.EventNameWalletBalanceLock, domain.EventWalletBalanceUpd))
}

func (w Wallet) unlockBalanceByUserId(ctx context.Context, delivery mbs.RPCDelivery) (resp *mbs.RPCResponse) {
	return w.walletUserIDRequest(ctx, delivery, func(ctx context.Context, req *domain.WalletUserIDRequest) (interface{}, error) {
		return w.walletUCase.UnlockByUserID(ctx, req.UserID)
	}, w.eventWallet(domain.EventNameWalletBalanceUnlock, domain.EventWalletBalanceUpd))
}

func (w Wallet) walletBalanceRequest(ctx context.Context, delivery mbs.RPCDelivery, fn func(context.Context, *domain.WalletBalanceRequest) (interface{}, error), events ...mbs.FnExecEvent) *mbs.RPCResponse {
	req := &domain.WalletBalanceRequest{}
	return w.rpc.Executor().Exec(ctx, req, delivery, func(ctx context.Context) (interface{}, error) {
		return fn(ctx, req)
	}, events...)
}

func (w Wallet) walletUserIDRequest(ctx context.Context, delivery mbs.RPCDelivery, fn func(context.Context, *domain.WalletUserIDRequest) (interface{}, error), events ...mbs.FnExecEvent) *mbs.RPCResponse {
	req := &domain.WalletUserIDRequest{}
	return w.rpc.Executor().Exec(ctx, req, delivery, func(ctx context.Context) (interface{}, error) {
		return fn(ctx, req)
	}, events...)
}

func (w Wallet) eventWallet(name domain.EventNameWallet, route domain.EventWallet) mbs.FnExecEvent {
	return func(bytes []byte) error {
		return w.pub.Publish(newEvent(name, route, bytes))
	}
}
