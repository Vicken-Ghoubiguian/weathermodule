package main

import "fmt"

type Weather struct {

}

type weatherTest struct {

	test string
}

func main() {

	weatherTest_1 := weatherTest{test: "Hello world !"}

	fmt.Printf("Weather test message: " + weatherTest_1.test + "\n")
}
