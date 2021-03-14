package web

import (
	"braydend/pi-door-opener/gpio"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sentryhttp "github.com/getsentry/sentry-go/http"
)

// RegisterRoutes - Register the HTTP endpoints that will be available
func RegisterRoutes() {
	handleAssets()
	handleIndex()
	handleToggleDoor()
	handleGetState()
}

func handleAssets() {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.Handle("/static/", sentryHandler.Handle(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))
}

func handleIndex() {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc("/", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}))
}

func handleToggleDoor() {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc("/toggle", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Toggling door state")
		gpio.SetPin(gpio.RelayPin, false)
		time.Sleep(5 * time.Second)
		gpio.SetPin(gpio.RelayPin, true)
	}))
}

func handleGetState() {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	type stateResponse struct {
		IsOpen bool
	}

	http.HandleFunc("/status", sentryHandler.HandleFunc(func(w http.ResponseWriter, t *http.Request) {
		var jsonResponse stateResponse
		if gpio.ReadPin(gpio.SensorPin) {
			jsonResponse = stateResponse{IsOpen: false}
		} else {
			jsonResponse = stateResponse{IsOpen: true}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}))
}
