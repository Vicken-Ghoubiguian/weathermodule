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
func (temp *Temperature) InitializeTemperature(value float64) {

	temp.temperatureValue = value
	temp.currentTemperatureScale = Kelvin
	temp.temperatureScaleSymbol = " K"
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

// Temperature conversion to Kelvin (°K)
func (temp *Temperature) SetTemperatureAsKelvin() {

	if temp.currentTemperatureScale == Celsius {

		temp.temperatureValue = temp.temperatureValue + 273.15
		temp.currentTemperatureScale = Kelvin
		temp.temperatureScaleSymbol = " K"

		fmt.Printf(Green() + "Temperature converted in Kelvin (K) successfully." + Reset() + "\n")

	} else if temp.currentTemperatureScale == Fahrenheit {

		temp.temperatureValue = (temp.temperatureValue - 32) * (5.0/9.0) + 273.15
		temp.currentTemperatureScale = Kelvin
		temp.temperatureScaleSymbol = " K"

		fmt.Printf(Green() + "Temperature converted in Kelvin (K) successfully." + Reset() + "\n")

	} else {

		fmt.Printf(Red() + "Temperature already in Kelvin (K)." + Reset() + "\n")

	}
}

// Temperature conversion to Celsius (°C)
func (temp *Temperature) SetTemperatureAsCelsius() {

	if temp.currentTemperatureScale == Kelvin {

		temp.temperatureValue = temp.temperatureValue - 273.15
		temp.currentTemperatureScale = Celsius
		temp.temperatureScaleSymbol = " °C"

		fmt.Printf(Green() + "Temperature converted in Celsius (°C) successfully." + Reset() + "\n")

	} else if temp.currentTemperatureScale == Fahrenheit {

		temp.temperatureValue = (temp.temperatureValue - 32) * (5.0/9.0)
		temp.currentTemperatureScale = Celsius
		temp.temperatureScaleSymbol = " °C"

		fmt.Printf(Green() + "Temperature converted in Celsius (°C) successfully." + Reset() + "\n")

	} else {

		fmt.Printf(Red() + "Temperature already in Celsius (°C)." + Reset() + "\n")

	}
}

// Temperature conversion to Fahrenheit (°F)
func (temp *Temperature) SetTemperatureAsFahrenheit() {

	if temp.currentTemperatureScale == Kelvin {

		temp.temperatureValue = (temp.temperatureValue - 273.15) * (9.0/5.0) + 32
		temp.currentTemperatureScale = Fahrenheit
                temp.temperatureScaleSymbol = " °F"

		fmt.Printf(Green() + "Temperature converted in Fahrenheit (°F) successfully." + Reset() + "\n")

	} else if temp.currentTemperatureScale == Celsius {

		temp.temperatureValue = (temp.temperatureValue * (9.0/5.0)) + 32
		temp.currentTemperatureScale = Fahrenheit
                temp.temperatureScaleSymbol = " °F"

		fmt.Printf(Green() + "Temperature converted in Fahrenheit (°F) successfully." + Reset() + "\n")

	} else {

		fmt.Printf(Red() + "Temperature already in Fahrenheit (°F)." + Reset() + "\n")

	}
}
