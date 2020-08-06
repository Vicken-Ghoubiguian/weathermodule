package main

import (
	"fmt"
	"weatherClasses"
)

//
type Weather struct {

	//
	coords weatherClasses.Coordinates
}

type weatherTest struct {

	test string
}

func main() {

	weatherTest_1 := weatherTest{test: "Hello world !"}
	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)

	fmt.Printf("Weather test message: " + weatherTest_1.test + "\n")
	fmt.Printf("(" + fmt.Sprintf("%f", coords.Longitude) + ", " + fmt.Sprintf("%f", coords.Latitude) + ")\n")
}
