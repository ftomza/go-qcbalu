package zapplus

import (
	"github.com/ftomza/go-qcbalu/pkg/viperplus"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func New(name string, config *viperplus.Viper) *Logger {

	logger := &zap.Logger{}
	var err error
	if config.IsProd() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	zapPlus := &Logger{logger.Named(name)}
	return zapPlus
}
