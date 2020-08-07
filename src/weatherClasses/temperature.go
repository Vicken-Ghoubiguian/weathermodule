package weatherClasses

//
import (
	"fmt"
)

//
type TemperatureScale int

//
const (

	Kelvin TemperatureScale = iota
	Celsius
	Fahrenheit
)

//
func (tempScale TemperatureScale) String() string {

	return [...]string{"Kelvin", "Celsius", "Fahrenheit"}[tempScale]
}

//
type Temperature struct {

	//
	temperatureValue float64
	currentTemperatureScale TemperatureScale
	temperatureScaleSymbol string
}

// Defining the Temperature initializer
func InitializeTemperature(value float64) *Temperature {

	return &Temperature{temperatureValue: value, currentTemperatureScale: Kelvin, temperatureScaleSymbol: " K"}
}

// 'temperatureValue' attribute getter
func (temp *Temperature) GetTemperatureValue() float64 {

	return temp.temperatureValue
}

// 'currentTemperatureScale' attribute getter
func (temp *Temperature) GetCurrentTemperatureScale() TemperatureScale {

	return temp.currentTemperatureScale
}

// 'temperatureScaleSymbol' attribute getter
func (temp *Temperature) GetTemperatureScaleSymbol() string {

	return temp.temperatureScaleSymbol
}

//
func (temp *Temperature) SetTemperatureAsKelvin() {

	if temp.currentTemperatureScale == Celsius {

		temp.temperatureValue = temp.temperatureValue + 273.15
		temp.currentTemperatureScale = Kelvin
		temp.temperatureScaleSymbol = " K"

		fmt.Printf("Temperature converted in Kelvin (K) successfully.\n")

	} else if temp.currentTemperatureScale == Fahrenheit {

		temp.temperatureValue = (temp.temperatureValue - 32) * (5/9) + 273.15
		temp.currentTemperatureScale = Kelvin
		temp.temperatureScaleSymbol = " K"

		fmt.Printf("Temperature converted in Kelvin (K) successfully.\n")

	} else {

		fmt.Printf("Temperature already in Kelvin (K).\n")

	}
}

//
func (temp *Temperature) SetTemperatureAsCelsius() {

	if temp.currentTemperatureScale == Kelvin {

		temp.temperatureValue = temp.temperatureValue - 273.15
		temp.currentTemperatureScale = Celsius
		temp.temperatureScaleSymbol = " °C"

		fmt.Printf("Temperature converted in Celsius (°C) successfully.\n")

	} else if temp.currentTemperatureScale == Fahrenheit {

		temp.temperatureValue = (temp.temperatureValue - 32) * (5/9)
		temp.currentTemperatureScale = Celsius
		temp.temperatureScaleSymbol = " °C"

		fmt.Printf("Temperature converted in Celsius (°C) successfully.\n")

	} else {

		fmt.Printf("Temperature already in Celsius (°C).\n")

	}
}

//

