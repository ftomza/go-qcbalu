// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ftomza/go-qcbalu/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// DeleteByUserID provides a mock function with given fields: ctx, userID
func (_m *WalletRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUserID provides a mock function with given fields: ctx, userID
func (_m *WalletRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	ret := _m.Called(ctx, userID)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Wallet); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: ctx
func (_m *WalletRepository) GetTransaction(ctx context.Context) domain.Transaction {
	ret := _m.Called(ctx)

	var r0 domain.Transaction
	if rf, ok := ret.Get(0).(func(context.Context) domain.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Transaction)
		}
	}

	return r0
}

// Store provides a mock function with given fields: ctx, item
func (_m *WalletRepository) Store(ctx context.Context, item *domain.Wallet) (*domain.Wallet, error) {
	ret := _m.Called(ctx, item)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Wallet) *domain.Wallet); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Wallet) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, item
func (_m *WalletRepository) Update(ctx context.Context, item *domain.Wallet) (*domain.Wallet, error) {
	ret := _m.Called(ctx, item)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Wallet) *domain.Wallet); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Wallet) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
