package ent

import (
	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/pkg/entplus"
)

// DomainWallet domain.Wallet object
type DomainWallet domain.Wallet

// ToWallet Convert to Wallet entity
func (w DomainWallet) ToWallet() *Wallet {
	item := &Wallet{}
	entplus.MustCopyValue(item, &w)
	return item
}

// ToDomainWallet Convert to domain.Wallet object
func (w Wallet) ToDomainWallet() *domain.Wallet {
	item := &domain.Wallet{}
	entplus.MustCopyValue(item, &w)
	return item
}

// ToDomainWallets Convert to domain.Wallet object list
func (w Wallets) ToDomainWallets() []*domain.Wallet {
	list := make([]*domain.Wallet, len(w))
	for i, item := range w {
		list[i] = item.ToDomainWallet()
	}
	return list
}
