package main

import (
	"fmt"
	event "listener/events"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer rabbitConn.Close()

	// start listening for messages
	log.Println("Listening and consuming RabbitMQ messages...")

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	count := 0

	// connect to RabbitMQ
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ is not ready yet...")

			// RabbitMQ usually takes some time in connecting
			time.Sleep(time.Second)
			count++
		} else {
			return c, nil
		}
		if count > 5 {
			return nil, err
		}
	}
}
