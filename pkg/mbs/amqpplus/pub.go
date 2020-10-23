package amqpplus

import (
	"fmt"

	"github.com/ftomza/go-qcbalu/pkg/mbs"

	"github.com/streadway/amqp"
)

type PUBService struct {
	ch       *amqp.Channel
	conn     *MBS
	exchange string
}

func (c *MBS) NewPUBService(exchange string) (svc mbs.PUBService) {

	svc = &PUBService{
		exchange: exchange,
		conn:     c,
	}
	return svc
}

func (s *PUBService) Publish(msg *mbs.PUBMessage) (err error) {
	if s.ch == nil {
		if s.ch, err = s.conn.createChannel("pub"); err != nil {
			return err
		}
	}
	if err = s.ch.Publish(
		s.exchange,
		msg.Route,
		false,
		false,
		amqp.Publishing{
			Headers:      amqp.Table(msg.Headers),
			ContentType:  ContentTypeBroker,
			Body:         msg.Data,
			DeliveryMode: amqp.Persistent,
		},
	); err != nil {
		return fmt.Errorf("ampqpplus/pub exchange publish : %s", err)
	}

	return nil
}
