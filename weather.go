package main

import "fmt"

type Weather struct {

	test string
}

func main() {

	weatherTest := Weather{test: "Hello world !"}

	fmt.Printf("Weather test message: " + weatherTest.test + "\n")
}
