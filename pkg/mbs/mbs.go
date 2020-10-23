package mbs

import (
	"context"
	"encoding/json"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/pkg/zapplus"
	"go.uber.org/zap"
)

var (
	RespErrBadRequest    = "BadRequest"
	RespErrInternalError = "InternalError"
)

var (
	HeaderEventName = "event"
)

type Headers map[string]interface{}

func NewEventHeaders(headers domain.EventHeaders) Headers {
	return Headers{
		HeaderEventName: headers.Event,
	}
}

type PUBMessage struct {
	Headers Headers
	Route   string
	Data    []byte
}

type RPCResponse struct {
	Error   *string
	Message []byte
}

func NewRPCResponseError(short string, err error) *RPCResponse {
	return &RPCResponse{
		Error:   &short,
		Message: []byte(err.Error()),
	}
}

func NewRPCResponse(data []byte) *RPCResponse {
	return &RPCResponse{
		Message: data,
	}
}

func (r *RPCResponse) Data() []byte {
	d, _ := json.Marshal(r)
	return d
}

type RPCMethod func(ctx context.Context, delivery RPCDelivery) (resp *RPCResponse)

type MBS interface {
	Close() error
}

type PUBService interface {
	Publish(msg *PUBMessage) (err error)
}

type RPCService interface {
	Run() (err error)
	AddMethod(path string, fn RPCMethod) RPCService
	Shutdown() error
	Executor() Executor
}

type Executor interface {
	Exec(ctx context.Context, req interface{}, delivery RPCDelivery, fn FnExec, events ...FnExecEvent) *RPCResponse
}

type RPCDelivery interface {
	Data() []byte
	ParseData(out interface{}) (err error)
}

type ResponseExecutor struct {
	Logger *zapplus.Logger
}

type FnExec func(context.Context) (interface{}, error)
type FnExecEvent func([]byte) error

func (e *ResponseExecutor) Exec(ctx context.Context, req interface{}, delivery RPCDelivery, fn FnExec, events ...FnExecEvent) *RPCResponse {
	err := delivery.ParseData(req)
	if err != nil {
		return NewRPCResponseError(RespErrBadRequest, err)
	}
	item, err := fn(ctx)
	if err != nil {
		return NewRPCResponseError(RespErrInternalError, err)
	}
	data, err := json.Marshal(item)
	if err != nil {
		return NewRPCResponseError(RespErrInternalError, err)
	}
	for _, ev := range events {
		if err := ev(data); err != nil {
			e.Logger.Error("rpc/executor: send event", zap.Error(err))
		}
	}
	return NewRPCResponse(data)
}
