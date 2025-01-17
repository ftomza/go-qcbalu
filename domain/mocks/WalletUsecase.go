// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ftomza/go-qcbalu/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WalletUsecase is an autogenerated mock type for the WalletUsecase type
type WalletUsecase struct {
	mock.Mock
}

// AddByUserID provides a mock function with given fields: ctx, userID, sum
func (_m *WalletUsecase) AddByUserID(ctx context.Context, userID uuid.UUID, sum int) (*domain.Wallet, error) {
	ret := _m.Called(ctx, userID, sum)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) *domain.Wallet); ok {
		r0 = rf(ctx, userID, sum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int) error); ok {
		r1 = rf(ctx, userID, sum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CredBalanceByUserID provides a mock function with given fields: ctx, userID, sum
func (_m *WalletUsecase) CredBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (*domain.Wallet, error) {
	ret := _m.Called(ctx, userID, sum)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) *domain.Wallet); ok {
		r0 = rf(ctx, userID, sum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int) error); ok {
		r1 = rf(ctx, userID, sum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DebtBalanceByUserID provides a mock function with given fields: ctx, userID, sum
func (_m *WalletUsecase) DebtBalanceByUserID(ctx context.Context, userID uuid.UUID, sum int) (*domain.Wallet, error) {
	ret := _m.Called(ctx, userID, sum)

	var r0 *domain.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) *domain.Wallet); ok {
		r0 = rf(ctx, userID, sum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int) error); ok {
		r1 = rf(ctx, userID, sum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelByUserID provides a mock function with given fields: ctx, userID
func (_m *WalletUsecase) DelByUserID(ctx context.Context, userID uuid.UUID) error {
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
func (_m *WalletUsecase) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
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

// LockByUserID provides a mock function with given fields: ctx, userID
func (_m *WalletUsecase) LockByUserID(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
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

// UnlockByUserID provides a mock function with given fields: ctx, userID
func (_m *WalletUsecase) UnlockByUserID(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
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
