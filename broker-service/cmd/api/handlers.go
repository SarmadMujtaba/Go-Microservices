package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Broker Hit!",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}
