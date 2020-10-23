package amqpplus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ftomza/go-qcbalu/pkg/mbs"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RPCService struct {
	executor *mbs.ResponseExecutor
	ch       *amqp.Channel
	conn     *MBS
	name     string
	queue    string
	done     chan error
	methods  map[string]mbs.RPCMethod
}

type RPCDelivery struct {
	delivery amqp.Delivery
}

func NewRPCDelivery(d amqp.Delivery) mbs.RPCDelivery {
	return &RPCDelivery{
		delivery: d,
	}
}

func (d RPCDelivery) Data() []byte {
	return d.delivery.Body
}

func (d RPCDelivery) ParseData(out interface{}) (err error) {
	return json.Unmarshal(d.Data(), out)
}

func (c *MBS) NewRPCService(name, queue string) (svc mbs.RPCService) {

	svc = &RPCService{
		name:  name,
		queue: queue,
		conn:  c,
		executor: &mbs.ResponseExecutor{
			Logger: c.logger,
		},
		methods: map[string]mbs.RPCMethod{},
		done:    make(chan error),
	}
	return svc
}

func (s *RPCService) Executor() mbs.Executor {
	return s.executor
}

func (s *RPCService) Run() (err error) {
	if s.ch == nil {
		s.ch, err = s.conn.createChannel("rpc/" + s.name)
		if err != nil {
			return err
		}
	}

	err = s.ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return fmt.Errorf("amqpplus/rpc: set qos: %w", err)
	}

	deliveries, err := s.ch.Consume(
		s.queue,
		s.name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("amqpplus/rpc: consume: %w", err)
	}

	go s.handle(deliveries)

	return nil
}

func (s *RPCService) AddMethod(path string, fn mbs.RPCMethod) mbs.RPCService {
	s.methods[path] = fn
	return s
}

func (s *RPCService) Shutdown() error {

	s.conn.logger.Info("amqpplus/rpc: Try shutdown consumer")

	if err := s.ch.Cancel(s.name, true); err != nil {
		return fmt.Errorf("amqpplus/rpc: consumer cancel failed: %s", err)
	}

	defer s.conn.logger.Info("amqpplus/rpc: shutdown OK", zap.String("name", s.name))

	return <-s.done
}

func (s *RPCService) handle(deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		ctx := context.Background()
		var err error
		s.conn.logger.Debug(fmt.Sprintf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		))

		fn, ok := s.methods[d.RoutingKey]
		resp := mbs.NewRPCResponseError("MethodNotFound", errors.New("rpc/handle: method not found: "+d.RoutingKey))
		if ok {
			resp = fn(ctx, NewRPCDelivery(d))
		}
		err = s.respond(d.CorrelationId, d.ReplyTo, resp.Data())
		if err != nil {
			s.conn.logger.Error("rpc/handle: Respond", zap.Error(err))
		}
		err = d.Ack(false)
		if err != nil {
			s.conn.logger.Error("rpc/handle: ACK", zap.Error(err))
		}
	}
	s.conn.logger.Info("rpc/handle: deliveries channel closed")
	s.done <- nil
}

func (s *RPCService) respond(correlationId, replayTo string, data []byte) error {
	return s.ch.Publish(
		"",
		replayTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   ContentTypeBroker,
			CorrelationId: correlationId,
			Body:          data,
			DeliveryMode:  amqp.Persistent,
		})
}
