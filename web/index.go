package web

import (
	"braydend/pi-door-opener/gpio"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterRoutes - Register the HTTP endpoints that will be available
func RegisterRoutes() {
	handleAssets()
	handleIndex()
	handleToggleDoor()
	handleGetState()
}

func handleAssets() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func handleIndex() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
}

func handleToggleDoor() {
	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Toggling door state")
		gpio.TogglePin(gpio.RelayPin)
	})
}

func handleGetState() {
	type stateResponse struct {
		IsOpen bool
	}

	http.HandleFunc("/status", func(w http.ResponseWriter, t *http.Request) {
		var jsonResponse stateResponse
		if gpio.ReadPin(gpio.SensorPin) {
			jsonResponse = stateResponse{IsOpen: true}
		} else {
			jsonResponse = stateResponse{IsOpen: false}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	})
}
