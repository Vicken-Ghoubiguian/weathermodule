package main

// Import all necessary packages
import (
	"fmt"
	"strconv"
	"weatherClasses"
)

// Defining the type 'Weather' which recover and manage current weather in a defined city
type Weather struct {

	//
	Coords weatherClasses.Coordinates

	//
	Temperature weatherClasses.Temperature
	FeelingLikeTemperature weatherClasses.Temperature
	MinTemperature weatherClasses.Temperature
	MaxTemperature weatherClasses.Temperature

	//
	UltraViolet weatherClasses.UV
}

// Defining the Weather initializer
func InitializeWeather() {

}

// main function to test all of the package
func main() {

	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(300.85)
	uv := weatherClasses.InitializeUV(10)

	fmt.Printf("(" + fmt.Sprintf("%f", coords.GetLongitude()) + ", " + fmt.Sprintf("%f", coords.GetLatitude()) + ")\n")

	temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + strconv.Itoa(uv.GetIndex()) + ", UV risk: " + uv.GetRisk().String() + "\n")
}
