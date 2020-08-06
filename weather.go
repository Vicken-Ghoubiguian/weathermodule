package main

import (
	"fmt"
	"weatherClasses"
)

//
type Weather struct {

	//
	Coords weatherClasses.Coordinates

	//
	Temperature weatherClasses.Temperature
}

type weatherTest struct {

	test string
}

func main() {

	weatherTest_1 := weatherTest{test: "Hello world !"}
	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(23.899)

	fmt.Printf("Weather test message: " + weatherTest_1.test + "\n")
	fmt.Printf("(" + fmt.Sprintf("%f", coords.Longitude) + ", " + fmt.Sprintf("%f", coords.Latitude) + ")\n")
	fmt.Printf("Temperature (in " + temperature.CurrentTemperatureScale + "): " + fmt.Sprintf("%f", temperature.TemperatureValue) + temperature.TemperatureScaleSymbol + "\n")
}
