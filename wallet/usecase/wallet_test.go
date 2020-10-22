package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/domain/mocks"
	"github.com/ftomza/go-qcbalu/wallet/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletUCaseTestSuite struct {
	suite.Suite
	Trans *mocks.Transaction
	Repo  *mocks.WalletRepository
	UCase domain.WalletUsecase
}

func (suite *WalletUCaseTestSuite) SetupTest() {
	suite.Repo = new(mocks.WalletRepository)
	suite.Trans = new(mocks.Transaction)
	suite.UCase, _ = usecase.NewWalletUsecase(suite.Repo, usecase.SetTimeout(time.Second*10))
}

func Test_EntUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(WalletUCaseTestSuite))
}

func (suite *WalletUCaseTestSuite) TearDownTest() {
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_AddByUserID() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(item, nil).Once()
		newItem, err := suite.UCase.AddByUserID(context.TODO(), item.ID, 0)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(nil, errors.New("fail")).Once()
		newItem, err := suite.UCase.AddByUserID(context.TODO(), item.ID, 0)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_DelByUserID() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("DeleteByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil).Once()
		err := suite.UCase.DelByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("DeleteByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(errors.New("fail")).Once()
		err := suite.UCase.DelByUserID(context.TODO(), item.UserID)
		suite.Error(err)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_GetByUserID() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(item, nil).Once()
		newItem, err := suite.UCase.GetByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("fail")).Once()
		newItem, err := suite.UCase.GetByUserID(context.TODO(), item.UserID)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_DebtBalanceByUserID() {
	suite.Run("ok", func() {
		item := *mocks.MockWallets["test1"]
		item.Balance = 30
		oldBalance := item.Balance
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.DebtBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.NoError(err)
		if suite.NotNil(newItem) {
			suite.Equal(newItem.Balance, oldBalance-10)
		}
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("fail")).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.DebtBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("lock", func() {
		item := *mocks.MockWallets["test1"]
		item.Balance = 30
		item.Lock = true
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.DebtBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.EqualError(err, "wallet: Item blocked")
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("insufficient", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.DebtBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.EqualError(err, "wallet: Insufficient funds")
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_CredBalanceByUserID() {
	suite.Run("ok", func() {
		item := *mocks.MockWallets["test1"]
		item.Balance = 30
		oldBalance := item.Balance
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.CredBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.NoError(err)
		if suite.NotNil(newItem) {
			suite.Equal(newItem.Balance, oldBalance+10)
		}
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("fail")).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.CredBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("lock", func() {
		item := *mocks.MockWallets["test1"]
		item.Balance = 30
		item.Lock = true
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.CredBalanceByUserID(context.TODO(), item.UserID, 10)
		suite.EqualError(err, "wallet: Item blocked")
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_LockByUserID() {
	suite.Run("ok", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.LockByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.True(newItem.Lock)
		suite.True(item.Lock)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("ok not update", func() {
		item := *mocks.MockWallets["test1"]
		item.Lock = true
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.LockByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.True(newItem.Lock)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("fail")).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.LockByUserID(context.TODO(), item.UserID)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}

func (suite *WalletUCaseTestSuite) Test_WalletUCase_UnlockByUserID() {
	suite.Run("ok", func() {
		item := *mocks.MockWallets["test1"]
		item.Lock = true
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Wallet")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.UnlockByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.False(newItem.Lock)
		suite.False(item.Lock)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("ok not update", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&item, nil).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Commit", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.UnlockByUserID(context.TODO(), item.UserID)
		suite.NoError(err)
		suite.NotNil(newItem)
		suite.False(newItem.Lock)
		suite.Repo.AssertExpectations(suite.T())
	})

	suite.Run("fail", func() {
		item := *mocks.MockWallets["test1"]
		suite.Repo.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("fail")).Once()
		suite.Repo.On("GetTransaction", mock.Anything).Return(suite.Trans).Once()
		suite.Trans.On("Begin", mock.Anything).Return(nil, nil).Once()
		suite.Trans.On("Rollback", mock.Anything, mock.Anything).Return(nil).Once()
		newItem, err := suite.UCase.UnlockByUserID(context.TODO(), item.UserID)
		suite.Error(err)
		suite.Nil(newItem)
		suite.Repo.AssertExpectations(suite.T())
	})
}
