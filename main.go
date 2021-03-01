package main

import (
	"braydend/pi-door-opener/gpio"
	"braydend/pi-door-opener/web"
	"log"
	"net/http"
)

func main() {
	pins := []gpio.PinConfig{
		{
			Number:  1,
			IsInput: true,
		},
		{
			Number:  gpio.RelayPin,
			IsInput: false,
		},
	}

	defer gpio.CloseGPIO()
	gpio.InitialiseGPIO(pins)
	web.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
