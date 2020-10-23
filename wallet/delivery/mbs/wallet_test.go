package mbs

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/ftomza/go-qcbalu/domain"

	"github.com/ftomza/go-qcbalu/pkg/mbs"
	"github.com/ftomza/go-qcbalu/pkg/zapplus"
	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"

	mocksDomain "github.com/ftomza/go-qcbalu/domain/mocks"
	"github.com/ftomza/go-qcbalu/pkg/mbs/mocks"
	"github.com/stretchr/testify/suite"
)

var (
	log = &zapplus.Logger{
		Logger: zap.NewExample().Named("TEST"),
	}
)

type WalletDeliveryTestSuite struct {
	suite.Suite
	Deliver *Wallet

	pub   *mocks.PUBService
	rpc   *mocks.RPCService
	uCase *mocksDomain.WalletUsecase
	dlv   *mocks.RPCDelivery
	ctx   context.Context
}

func (suite *WalletDeliveryTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.pub = new(mocks.PUBService)
	suite.rpc = new(mocks.RPCService)
	suite.dlv = new(mocks.RPCDelivery)
	suite.uCase = new(mocksDomain.WalletUsecase)

	suite.rpc.On("AddMethod",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("mbs.RPCMethod")).Return(suite.rpc).Times(7)

	suite.rpc.On("Executor").Return(&mbs.ResponseExecutor{Logger: log}).Once()

	suite.Deliver = NewMBSWallet(suite.rpc, suite.pub, suite.uCase)

}

func dataX(t *testing.T, in interface{}) []byte {
	d, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}
	return d
}

func Test_EntUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(WalletDeliveryTestSuite))
}

func (suite *WalletDeliveryTestSuite) TearDownTest() {
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_addBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(nil).Once()
		suite.uCase.On("AddByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("int")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.addBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "self"},
			Route:   "balance.add",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})

	suite.Run("fail fn", func() {
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(nil).Once()
		suite.uCase.On("AddByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("int")).Return(nil, errors.New("fail")).Once()
		suite.rpc.On("Executor").Return(&mbs.ResponseExecutor{Logger: log}).Once()
		resp := suite.Deliver.addBalanceByUserId(suite.ctx, suite.dlv)
		if suite.NotNil(resp.Error) {
			suite.Equal(*resp.Error, "InternalError")
			suite.Equal(resp.Message, []byte("fail"))
		}
	})

	suite.Run("fail parse", func() {
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(errors.New("fail")).Once()
		suite.rpc.On("Executor").Return(&mbs.ResponseExecutor{Logger: log}).Once()
		resp := suite.Deliver.addBalanceByUserId(suite.ctx, suite.dlv)
		if suite.NotNil(resp.Error) {
			suite.Equal(*resp.Error, "BadRequest")
			suite.Equal(resp.Message, []byte("fail"))
		}
	})

	suite.Run("fail event", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(nil).Once()
		suite.uCase.On("AddByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("int")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(errors.New("fail")).Once()
		suite.rpc.On("Executor").Return(&mbs.ResponseExecutor{Logger: log}).Once()
		resp := suite.Deliver.addBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "self"},
			Route:   "balance.add",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_delBalanceByUserId() {
	suite.Run("ok", func() {
		data := dataX(suite.T(), &domain.Wallet{})
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletUserIDRequest")).Return(nil).Once()
		suite.uCase.On("DelByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.delBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "self"},
			Route:   "balance.del",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_getBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletUserIDRequest")).Return(nil).Once()
		suite.uCase.On("GetByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(item, nil).Once()
		resp := suite.Deliver.getBalanceByUserId(suite.ctx, suite.dlv)
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_debtBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(nil).Once()
		suite.uCase.On("DebtBalanceByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("int")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.debtBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "debt"},
			Route:   "balance.upd",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_credBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletBalanceRequest")).Return(nil).Once()
		suite.uCase.On("CredBalanceByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("int")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.credBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "cred"},
			Route:   "balance.upd",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_lockBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletUserIDRequest")).Return(nil).Once()
		suite.uCase.On("LockByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.lockBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "lock"},
			Route:   "balance.upd",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}

func (suite *WalletDeliveryTestSuite) Test_WalletDelivery_unlockBalanceByUserId() {
	suite.Run("ok", func() {
		item := mocksDomain.MockWallets["test1"]
		data := dataX(suite.T(), item)
		suite.dlv.On("ParseData", mock.AnythingOfType("*domain.WalletUserIDRequest")).Return(nil).Once()
		suite.uCase.On("UnlockByUserID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(item, nil).Once()
		suite.pub.On("Publish", mock.AnythingOfType("*mbs.PUBMessage")).Return(nil).Once()
		resp := suite.Deliver.unlockBalanceByUserId(suite.ctx, suite.dlv)
		suite.pub.AssertCalled(suite.T(), "Publish", &mbs.PUBMessage{
			Headers: mbs.Headers{"event": "unlock"},
			Route:   "balance.upd",
			Data:    data,
		})
		suite.Empty(resp.Error)
		suite.Equal(resp.Message, data)
	})
}
