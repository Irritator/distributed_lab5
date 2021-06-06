package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("test_direct", false, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	infiniteChanel := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
		}
	}()
	<-infiniteChanel
}
