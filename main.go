package main

import (
	"braydend/pi-door-opener/gpio"
	"braydend/pi-door-opener/web"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func configureLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func init() {
	configureLogger()
}

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

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              "https://ca703ac80a0b41ce80a6f5189af6f4d0@o538041.ingest.sentry.io/5655995",
		Environment:      os.Getenv("ENV"),
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	defer gpio.CloseGPIO()
	gpio.InitialiseGPIO(pins)
	web.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
