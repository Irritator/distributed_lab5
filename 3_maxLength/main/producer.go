package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("limitedQueue", false, false, false, false, amqp.Table{"x-max-length": 10, "x-overflow": "reject-publish"})
	for i := 0; i < 20; i++ {
		body := "test message" + strconv.Itoa(i)
		err := ch.Publish(
			"", q.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(body)
		}
	}
}
