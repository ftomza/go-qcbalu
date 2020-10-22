package mocks

import (
	"github.com/ftomza/go-qcbalu/domain"
	"github.com/google/uuid"
)

var (
	MockWallets = map[string]*domain.Wallet{
		"test1": {ID: uuid.New(), UserID: uuid.New(), Lock: false, Balance: 0},
		"test2": {ID: uuid.New(), UserID: uuid.New(), Lock: false, Balance: 0},
		"test3": {ID: uuid.New(), UserID: uuid.New(), Lock: false, Balance: 0},
	}
)
