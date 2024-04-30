package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = "3001"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Listening on Port: %s", Port)

	// Define Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}

	// Run Server
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
