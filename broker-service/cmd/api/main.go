package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const Port = "3001"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	// connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer rabbitConn.Close()
	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Broker - Listening on Port: %s", Port)

	// Define Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}

	// Run Server
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
