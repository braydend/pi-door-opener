package gpio

import (
	"log"

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
		panic(err)
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
		panic(err)
	}
}

// SetPin - Set a pin to a specific high/low value
func SetPin(pinNumber uint, high bool) {
	pin := rpio.Pin(pinNumber)

	if high {
		pin.High()
	} else {
		pin.Low()
	}
}

// TogglePin - Toggle specified pin Low->High->Low
func TogglePin(pinNumber uint) {
	pin := rpio.Pin(pinNumber)

	pin.Toggle()
}

// ReadPin - Read state of specified pin
func ReadPin(pinNumber uint) bool {
	pin := rpio.Pin(pinNumber)

	if pin.Read() == rpio.High {
		return true
	}
	return false
}
