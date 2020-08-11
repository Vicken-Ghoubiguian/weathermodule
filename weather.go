package main

// Import all necessary packages
import (
	"fmt"
	"strconv"
	"weatherClasses"
	"io/ioutil"
	"net/http"
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
func InitializeWeather(city string, apiKey string) (string, string) {

	//
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	uvRequest := fmt.Sprintf("http://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, "48.85", "2.3")

	//
	weatherResp, _ := http.Get(weatherRequest)
	uvResp, _ := http.Get(uvRequest)

	//
	weatherJsonString, _ := ioutil.ReadAll(weatherResp.Body)
	uvJsonString, _ := ioutil.ReadAll(uvResp.Body)

	//
	return string(weatherJsonString), string(uvJsonString)
}

// main function to test all of the package
func main() {

	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(300.85)
	uv := weatherClasses.InitializeUV(10)
	weatherResponse, uvResponse := InitializeWeather("Paris,Fr", "5222a1c311ca31001b0877137d584c36")

	fmt.Printf("Weather response: " + weatherResponse + "\n\n")
	fmt.Printf("UV response: " + uvResponse + "\n\n")

	fmt.Printf("(" + fmt.Sprintf("%f", coords.GetLongitude()) + ", " + fmt.Sprintf("%f", coords.GetLatitude()) + ")\n")

	temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + strconv.Itoa(uv.GetIndex()) + ", UV risk: " + uv.GetRisk().String() + "\n")
}
