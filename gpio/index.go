package gpio

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

// RelayPin - Pin for GPIO output controlling relay
const RelayPin = 2

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

	if err != nil {
		panic(err)
	}
}

// TogglePin - Toggle specified pin Low->High->Low
func TogglePin(pinNumber uint) {
	pin := rpio.Pin(pinNumber)

	pin.Toggle()
}
