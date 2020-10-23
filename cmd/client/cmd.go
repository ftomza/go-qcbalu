package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"

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

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	reader := bufio.NewReader(os.Stdin)
	for {
		log.Print("[>] Command: ")
		text, _ := reader.ReadString('\n')
		log.Printf("[+] Try: %s", text)
		parts := strings.SplitN(strings.TrimSpace(text), " ", 2)
		key := parts[0]
		if key == "quit" {
			return
		}
		if key == "" {
			continue
		}
		body := parts[1]
		corrId := uuid.New().String()
		err = ch.Publish(
			"rpc.wallet",
			key,
			false,
			false,
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: corrId,
				ReplyTo:       q.Name,
				Body:          []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		log.Printf("[x] Requesting %s(%s)", key, body)

		for d := range msgs {
			if corrId == d.CorrelationId {
				in := struct {
					Error   string
					Message []byte
				}{}
				err := json.Unmarshal(d.Body, &in)
				failOnError(err, fmt.Sprintf("Failed parse body: %v", d.Body))
				log.Printf("[.] Got Error: %s, Message: %s", in.Error, string(in.Message))
				break
			}
		}
	}

}

func init() {
	flag.Parse()
}
