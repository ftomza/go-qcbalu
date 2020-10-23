package amqpplus

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/ftomza/go-qcbalu/pkg/zapplus"

	"github.com/streadway/amqp"
)

var (
	RespErrBadRequest    = "BadRequest"
	RespErrInternalError = "InternalError"
)

var (
	ContentTypeBroker = "application/json"
)

type options struct {
	timeout time.Duration
}

type Option func(*options)

func SetTimeout(duration time.Duration) Option {
	return func(o *options) {
		o.timeout = duration
	}
}

type MBS struct {
	cn     *amqp.Connection
	logger *zapplus.Logger
	options
}

func NewAMQPPlusConn(cn *amqp.Connection, logger *zapplus.Logger, opts ...Option) (*MBS, error) {
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}

	if cn == nil {
		return nil, errors.New("amqpplus: connections not set")
	}

	return &MBS{
		cn:      cn,
		logger:  logger,
		options: options,
	}, nil
}

func (c *MBS) Close() (err error) {
	return c.cn.Close()
}

func (c *MBS) createChannel(name string) (ch *amqp.Channel, err error) {
	if ch, err = c.cn.Channel(); err != nil {
		return nil, fmt.Errorf("ampqpplus: create channel '%s': %w", name, err)
	}
	go c.notifyClosing(name)
	return
}

func (c *MBS) notifyClosing(name string) {
	err := <-c.cn.NotifyClose(make(chan *amqp.Error))
	msg := "amqpplus: closing connection"
	zString := zap.String("channel", name)
	if err != nil {
		c.logger.Error(msg, zap.Error(err), zString)
	} else {
		c.logger.Info(msg, zString)
	}

}
