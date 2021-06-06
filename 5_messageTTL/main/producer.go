package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("test_ttl", false, false, false, false, amqp.Table{"x-message-ttl": 60000})
	for i := 0; i < 10; i++ {
		body := "test message" + strconv.Itoa(i)
		_ = ch.Publish(
			"", q.Name, false, false,
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(body),
				DeliveryMode: 2,
				Expiration:   "20000",
			})
		fmt.Println(body)
	}
}
