package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	_ = ch.ExchangeDeclare("test_subscribe", "fanout", true, false, false, false, nil)
	for i := 0; i < 10; i++ {
		body := "test message" + strconv.Itoa(i)
		_ = ch.Publish("test_subscribe", "", false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		fmt.Println(body + " sent!")
	}
}
