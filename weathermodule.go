package main

// Import all necessary packages
import (
	"fmt"
	"os"
	"strconv"
	"weatherClasses"
	"io/ioutil"
	"net/http"
)

// Importation of the github project gjson to treat received json
//import "github.com/tidwall/gjson"

var red string = "\033[31m"
var green string = "\033[32m"
var cyan string = "\033[36m"
var reset string = "\033[0m"

// Function which display HTTP request error's code and message when the first occurs
func owmErrorHandler(codeError string, errorMessage string) {

	fmt.Println(red + "Occured error (" + codeError + "): " + errorMessage + reset)

	fmt.Println("\n")

	os.Exit(1)
}

// Function which display other errors when they occurs
func otherErrorHandlerFunction(err error) {

	if err != nil {

		fmt.Println(red + err.Error() + reset)

		os.Exit(1)
	}
}

// Defining the type 'Weather' which recover and manage current weather in a defined city
type WeatherModule struct {

	//
	Coords weatherClasses.Coordinates

	//
	Weather weatherClasses.Weather

	//
	Temperature weatherClasses.Temperature
	FeelingLikeTemperature weatherClasses.Temperature
	MinTemperature weatherClasses.Temperature
	MaxTemperature weatherClasses.Temperature

	//
	Sunrise weatherClasses.SunTime
	Sunset weatherClasses.SunTime

	//
	UltraViolet weatherClasses.UV
}

// Defining the Weather initializer
func InitializeWeatherModule(city string, apiKey string) (string, string) {

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	uvRequest := fmt.Sprintf("http://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, "48.85", "2.3")

	//
	weatherResp, err0 := http.Get(weatherRequest)

	//
	otherErrorHandlerFunction(err0)

	//
	uvResp, err1 := http.Get(uvRequest)

	//
	otherErrorHandlerFunction(err1)

	//
	weatherJsonString, err2 := ioutil.ReadAll(weatherResp.Body)

	//
	otherErrorHandlerFunction(err2)

	//
	uvJsonString, err3 := ioutil.ReadAll(uvResp.Body)

	//
	otherErrorHandlerFunction(err3)

	return string(weatherJsonString), string(uvJsonString)
}

// main function to test all of the package
func main() {

	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(300.85)
	uv := weatherClasses.InitializeUV(10)
	sunRise := weatherClasses.InitializeSunTime(1597221424)
	sunSet := weatherClasses.InitializeSunTime(1597221424)
	weather := weatherClasses.InitializeWeather(300, "Drizzle", "light intensity drizzle", "09d")
	weatherResponse, uvResponse := InitializeWeatherModule("Paris,Fr", "")

	fmt.Printf("Weather response: " + weatherResponse + "\n\n")
	fmt.Printf("UV response: " + uvResponse + "\n\n")

	fmt.Printf("Weather (" + weather.GetMain() + ", " + weather.GetDescription() + ", " + weather.GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", coords.GetLongitude()) + ", " + fmt.Sprintf("%f", coords.GetLatitude()) + ")\n")

	temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", temperature.GetTemperatureValue()) + temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + strconv.Itoa(uv.GetIndex()) + ", UV risk: " + uv.GetRisk().String() + "\n")

	sunRise.SetCurrentFormatAsTimestamp()
	sunSet.SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + sunRise.GetCurrentFormat().String() + "): " + sunRise.GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + sunSet.GetCurrentFormat().String() + "): " + sunSet.GetSunTimeInCurrentFormat() + "\n")
}
