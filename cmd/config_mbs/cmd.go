package main

import (
	"flag"
	"log"

	"github.com/streadway/amqp"
)

var (
	uri = flag.String("uri", "amqp://localhost/qcbalu", "AMQP URI")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	cn, err := amqp.Dial(*uri)
	failOnError(err, "Dial")
	defer cn.Close()
	ch, err := cn.Channel()
	failOnError(err, "Channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"rpc.wallet",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a Exchange rpc.wallet")

	err = ch.ExchangeDeclare(
		"pub.wallet",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a Exchange pub.wallet")

	q, err := ch.QueueDeclare(
		"rpc.wallet.main",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue rpc.wallet.main")

	err = ch.QueueBind(
		q.Name,
		"*",
		"rpc.wallet",
		false,
		nil,
	)
	failOnError(err, "Failed bind rpc.wallet.main -> rpc.wallet")

	q, err = ch.QueueDeclare(
		"pub.wallet.main.client1",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue rpc.wallet.main key *")

	err = ch.QueueBind(
		q.Name,
		"balance.*",
		"pub.wallet",
		false,
		nil,
	)
	failOnError(err, "Failed bind rpc.wallet.main -> pub.wallet key balance.*")
}

func init() {
	flag.Parse()
}
