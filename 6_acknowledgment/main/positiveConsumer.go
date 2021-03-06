package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("test_ack", false, false, false, false, amqp.Table{"x-message-ttl": 60000})
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

	infiniteChanel := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
			_ = d.Ack(false)
		}
	}()
	<-infiniteChanel
}
