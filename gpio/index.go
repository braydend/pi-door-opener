package gpio

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/stianeikeland/go-rpio/v4"
)

// RelayPin - Pin for GPIO output controlling relay
const RelayPin = 2

// SensorPin - Pin for GPIO input reading magnet switch
const SensorPin = 3

// PinConfig - Map a GPIO pin with an input/output mode
type PinConfig struct {
	Number  uint
	IsInput bool
}

// InitialiseGPIO - Set up GPIO
func InitialiseGPIO(config []PinConfig) {
	log.Println("Attempting to initialise GPIO")
	err := rpio.Open()

	if err != nil {
		sentry.CaptureException(err)
		log.Panic(err)
	}

	configureGPIO(config)
}

func configureGPIO(config []PinConfig) {
	for _, pinConfig := range config {
		pin := rpio.Pin(pinConfig.Number)
		var mode string
		if pinConfig.IsInput {
			mode = "Input"
			pin.Input()
		} else {
			mode = "Output"
			pin.Output()
			if (os.Getenv("INITIALISE_HIGH") == "true") {
				pin.High()
			}
		}
		log.Printf("Setting pin %d to %s.\n", pinConfig.Number, mode)
	}
}

//CloseGPIO - Unmap GPIO
func CloseGPIO() {
	log.Println("Attempting to clean up GPIO")
	err := rpio.Close()

	SetPin(RelayPin, true)
	if err != nil {
		sentry.CaptureException(err)
		log.Panic(err)
	}
}

// SetPin - Set a pin to a specific high/low value
func SetPin(pinNumber uint, high bool) {
	pin := rpio.Pin(pinNumber)

	if high {
		log.Printf("Setting pin %d to high\n", pinNumber)

		pin.High()
	} else {
		log.Printf("Setting pin %d to low\n", pinNumber)

		pin.Low()
	}
}

// TogglePin - Toggle specified pin Low->High->Low
func TogglePin(pinNumber uint) {
	pin := rpio.Pin(pinNumber)
	log.Printf("Toggling pin %d\n", pinNumber)

	pin.Toggle()
}

// ReadPin - Read state of specified pin
func ReadPin(pinNumber uint) bool {
	pin := rpio.Pin(pinNumber)
	state := pin.Read()
	isHigh := state == rpio.High
	log.Printf("Pin %d read as %d\n", pinNumber, state)

	if os.Getenv("INVERT_DOOR_SENSOR") == "true" {
		return !isHigh
	}
	return isHigh
}
