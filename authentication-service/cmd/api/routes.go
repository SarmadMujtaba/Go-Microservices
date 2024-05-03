package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Config) routes() http.Handler {
	route := mux.NewRouter()

	// route.HandleFunc("/", app.Broker).Methods(http.MethodPost)

	http.ListenAndServe(":"+Port, route)
	return route
}
