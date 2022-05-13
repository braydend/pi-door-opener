package web

import (
	"braydend/pi-door-opener/gpio"
	"encoding/json"
	"fmt"
	"log"
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

func logRoute(description, route string) {
	log.Printf("Registering %s route: %s\n", description, route)
}

func handleAssets() {
	route := "/static/"
	logRoute("static assets", route)
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.Handle(route, sentryHandler.Handle(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))
}

func handleIndex() {
	route := "/"
	logRoute("index", route)
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc(route, sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}))
}

func handleToggleDoor() {
	route := "/toggle"
	logRoute("toggle door", route)
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc(route, sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Toggling door state. Request from: %s", r.Host)
		fmt.Fprintf(w, "Toggling door state")
		gpio.SetPin(gpio.RelayPin, false)
		time.Sleep(5 * time.Second)
		gpio.SetPin(gpio.RelayPin, true)
	}))
}

func handleGetState() {
	route := "/status"
	logRoute("current status", route)
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	type stateResponse struct {
		IsOpen bool
	}

	http.HandleFunc(route, sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonResponse stateResponse
		if gpio.ReadPin(gpio.SensorPin) {
			jsonResponse = stateResponse{IsOpen: false}
		} else {
			jsonResponse = stateResponse{IsOpen: true}
		}
		log.Printf("Reading door state(%v). Request from: %s", jsonResponse, r.Host)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}))
}
