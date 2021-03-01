package web

import (
	"braydend/pi-door-opener/gpio"
	"fmt"
	"net/http"
)

// RegisterRoutes - Register the HTTP endpoints that will be available
func RegisterRoutes() {
	handleToggleDoor()
	handleGetState()
}

func handleToggleDoor() {
	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Toggling door state")
		gpio.TogglePin(2)
	})
}

func handleGetState() {
	http.HandleFunc("/status", func(w http.ResponseWriter, t *http.Request) {
		fmt.Fprintf(w, "I am the current state")
	})
}
