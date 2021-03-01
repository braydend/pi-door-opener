package main

import (
	"braydend/pi-door-opener/gpio"
	"braydend/pi-door-opener/web"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	pins := []gpio.PinConfig{
		{
			Number:  gpio.SensorPin,
			IsInput: true,
		},
		{
			Number:  gpio.RelayPin,
			IsInput: false,
		},
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://ca703ac80a0b41ce80a6f5189af6f4d0@o538041.ingest.sentry.io/5655995",
		Environment:      "dev",
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	defer gpio.CloseGPIO()
	gpio.InitialiseGPIO(pins)
	web.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
