package main

import (
	"net/http"
)

func (app *Config) Broker(rw http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	_ = app.writeJSON(rw, http.StatusOK, payload)

}
