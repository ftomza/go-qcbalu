package repository_test

import (
	"context"
	"testing"

	"github.com/ftomza/go-qcbalu/domain/mocks"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/wallet"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/wallet/repository"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent/enttest"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type EntWalletRepositoryTestSuite struct {
	suite.Suite
	Ctx    context.Context
	Client *ent.Client
	Repo   domain.WalletRepository
}

func (suite *EntWalletRepositoryTestSuite) SetupTest() {
	suite.Client = enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1").Debug()
	suite.Ctx = context.Background()
	suite.Repo = repository.NewEntWalletRepository(suite.Client)
}

func Test_EntWalletRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EntWalletRepositoryTestSuite))
}

func (suite *EntWalletRepositoryTestSuite) TearDownTest() {
	_ = suite.Client.Close()
}

func (suite *EntWalletRepositoryTestSuite) Test_EntWalletRepository_Store() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		_, err := suite.Repo.Store(suite.Ctx, item)
		suite.NoError(err, "Store")
		obj, err := suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
		suite.NoError(err, "Query")
		suite.Equal(item.Lock, obj.Lock)
		suite.Equal(item.ID, obj.ID)
		suite.Equal(item.Balance, obj.Balance)
		suite.NotEmpty(obj.CreateTime)
		suite.NotEmpty(obj.UpdateTime)
		suite.NotEmpty(obj.Version)
	})

	suite.Run("wrong balance", func() {
		item := *mocks.MockWallets["test2"]
		item.Balance = -20
		_, err := suite.Repo.Store(suite.Ctx, &item)
		suite.EqualError(err, "ent: validator failed for field \"balance\": value out of range")
		_, err = suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
		suite.EqualError(err, "ent: wallet not found")
	})

	suite.Run("nil item", func() {
		_, err := suite.Repo.Store(suite.Ctx, nil)
		suite.EqualError(err, "repo: Nil item")
	})
}

func (suite *EntWalletRepositoryTestSuite) Test_EntWalletRepository_Update() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		userForUpdate, err := suite.Repo.Store(suite.Ctx, item)
		if suite.NoError(err, "Store") {
			userForUpdate.Balance = 10
			_, err = suite.Repo.Update(suite.Ctx, userForUpdate)
			if suite.NoError(err, "Update") {
				obj, err := suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
				if suite.NoError(err, "Query") {
					suite.Equal(userForUpdate.Balance, obj.Balance)
					suite.NotEqual(userForUpdate.Version, obj.Version)
				}
			}
		}
	})

	suite.Run("wrong version", func() {
		item := mocks.MockWallets["test2"]
		_, err := suite.Repo.Store(suite.Ctx, item)
		suite.NoError(err, "Store")
		userForUpdate, err := suite.Repo.GetByUserID(suite.Ctx, item.UserID)
		if suite.NoError(err, "Get") {
			suite.Client.Wallet.UpdateOneID(item.ID).SetBalance(60).SetVersion(userForUpdate.Version).ExecX(suite.Ctx)
			userForUpdate.Balance = 20
			_, err = suite.Repo.Update(suite.Ctx, userForUpdate)
			suite.EqualError(err, "ent: version not valid")
		}
	})

	suite.Run("nil item", func() {
		_, err := suite.Repo.Update(suite.Ctx, nil)
		suite.EqualError(err, "repo: Nil item")
	})
}

func (suite *EntWalletRepositoryTestSuite) Test_EntWalletRepository_GetByUserID() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		_, err := suite.Repo.Store(suite.Ctx, item)
		suite.NoError(err, "Store")
		userOnStore, err := suite.Repo.GetByUserID(suite.Ctx, item.UserID)
		if suite.NoError(err, "Get") {
			obj, err := suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
			if suite.NoError(err, "Query") {
				suite.Equal(userOnStore.Balance, obj.Balance)
				suite.Equal(userOnStore.Version, obj.Version)
				suite.Equal(userOnStore.CreateTime, obj.CreateTime)
				suite.Equal(userOnStore.UserID, obj.UserID)
				suite.Equal(userOnStore.UpdateTime, obj.UpdateTime)
				suite.Equal(userOnStore.Lock, obj.Lock)
			}
		}
	})

	suite.Run("not found", func() {
		item := mocks.MockWallets["test2"]
		userForUpdate, err := suite.Repo.GetByUserID(suite.Ctx, item.UserID)
		suite.Nil(userForUpdate)
		suite.EqualError(err, "ent: wallet not found")
	})
}

func (suite *EntWalletRepositoryTestSuite) Test_EntWalletRepository_DeleteByUserID() {
	suite.Run("ok", func() {
		item := mocks.MockWallets["test1"]
		_, err := suite.Repo.Store(suite.Ctx, item)
		suite.NoError(err, "Store")
		err = suite.Repo.DeleteByUserID(suite.Ctx, item.UserID)
		if suite.NoError(err, "Delete") {
			_, err = suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
			suite.EqualError(err, "ent: wallet not found")
		}
	})

	suite.Run("not found ok", func() {
		item := mocks.MockWallets["test2"]
		err := suite.Repo.DeleteByUserID(suite.Ctx, item.UserID)
		if suite.NoError(err, "Delete") {
			_, err = suite.Client.Wallet.Query().Where(wallet.UserID(item.UserID)).Only(suite.Ctx)
			suite.EqualError(err, "ent: wallet not found")
		}
	})
}
