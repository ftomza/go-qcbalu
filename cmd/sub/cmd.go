package main

import (
	"flag"
	"log"

	"github.com/streadway/amqp"
)

var (
	uri  = flag.String("uri", "amqp://localhost/qcbalu", "AMQP URI")
	name = flag.String("name", "client1", "Client name")
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

	msgs, err := ch.Consume(
		"pub.wallet.main."+*name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Headers: %v, Route: %s -> %s", d.Headers, d.RoutingKey, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for events. To exit press CTRL+C")
	<-forever
}

func init() {
	flag.Parse()
}
