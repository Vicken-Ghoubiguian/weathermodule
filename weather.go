package main

// Import all necessary packages
import (
	"fmt"
	"weatherClasses"
)

// Defining the type 'Weather' which recover and manage current weather in a defined city
type Weather struct {

	//
	Coords weatherClasses.Coordinates

	//
	Temperature weatherClasses.Temperature
}

// Defining 'weatherTest' for a test struct
type weatherTest struct {

	test string
}

// main function to test all of the package
func main() {

	weatherTest_1 := weatherTest{test: "Hello world !"}
	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(23.899)

	fmt.Printf("Weather test message: " + weatherTest_1.test + "\n")
	fmt.Printf("(" + fmt.Sprintf("%f", coords.GetLongitude()) + ", " + fmt.Sprintf("%f", coords.GetLatitude()) + ")\n")
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")
}
