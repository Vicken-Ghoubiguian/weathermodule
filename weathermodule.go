package main

// Import all necessary packages
import (
	"fmt"
	"os"
	"strconv"
	"weatherClasses"
	"io/ioutil"
	"net/http"
	"strings"
)

// Importation of the github project gjson to treat received json
import "github.com/tidwall/gjson"

var red string = "\033[31m"
var green string = "\033[32m"
var cyan string = "\033[36m"
var reset string = "\033[0m"

// Function which extracts weather datas from JSON response
func extractWeatherFromJSONFunction(weatherFromHTTPResponseString string) string {

	brutWeatherWithoutHooks := strings.Trim(weatherFromHTTPResponseString, "[]")

	return brutWeatherWithoutHooks
}

// Function which display HTTP request error's code and message when the first occurs
func owmErrorHandlerFunction(codeError string, errorMessage string) {

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
	Coords *weatherClasses.Coordinates

	//
	Weather *weatherClasses.Weather

	//
	Temperature *weatherClasses.Temperature
	FeelingLikeTemperature *weatherClasses.Temperature
	MinTemperature *weatherClasses.Temperature
	MaxTemperature *weatherClasses.Temperature

	//
	Sunrise *weatherClasses.SunTime
	Sunset *weatherClasses.SunTime

	//
	UltraViolet *weatherClasses.UV
}

// Defining the Weather initializer
func InitializeWeatherModule(city string, apiKey string) *WeatherModule {

	//
	var currentCoordinates *weatherClasses.Coordinates
	var currentWeather *weatherClasses.Weather
	var currentTemperature *weatherClasses.Temperature
	var currentFeelingTemperature *weatherClasses.Temperature
	var currentMinimumTemperature *weatherClasses.Temperature
	var currentMaximumTemperature *weatherClasses.Temperature
	var currentSunrise *weatherClasses.SunTime
	var currentSunset *weatherClasses.SunTime
	var currentUV *weatherClasses.UV

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	//
	weatherResp, err0 := http.Get(weatherRequest)

	//
	otherErrorHandlerFunction(err0)

	//
	weatherJsonString, err2 := ioutil.ReadAll(weatherResp.Body)

	//
	otherErrorHandlerFunction(err2)

	//
	owmCode := gjson.Get(string(weatherJsonString), "cod")

	//
	if owmCode.Int() != 200 {

		//
		owmMessage := gjson.Get(string(weatherJsonString), "message")

		// Calling the 'owmErrorHandlerFunction' to treat the current error...
		owmErrorHandlerFunction(owmCode.String(), owmMessage.String())

	} else {

		//
		weather := extractWeatherFromJSONFunction(gjson.Get(string(weatherJsonString), "weather").String())

		//
		currentCoordinates = weatherClasses.InitializeCoordinates(gjson.Get(string(weatherJsonString), "coord.lon").Float(), gjson.Get(string(weatherJsonString), "coord.lat").Float())

		//
		currentWeather = weatherClasses.InitializeWeather(gjson.Get(weather, "id").Int(), gjson.Get(weather, "main").String(), gjson.Get(weather, "description").String(), gjson.Get(weather, "icon").String())

		//
		currentTemperature = weatherClasses.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp").Float())
		currentFeelingTemperature = weatherClasses.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.feels_like").Float())
		currentMinimumTemperature = weatherClasses.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_min").Float())
		currentMaximumTemperature = weatherClasses.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_max").Float())

		//
		currentSunrise = weatherClasses.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunrise").Int())
		currentSunset = weatherClasses.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunset").Int())

		//
		currentUV = weatherClasses.InitializeUV(10)

		// Displaying success message...
		fmt.Println(green + "Weather implemented successfully !" + reset + "\n")
	}

	return &WeatherModule{Coords: currentCoordinates, Weather: currentWeather, Temperature: currentTemperature, FeelingLikeTemperature: currentFeelingTemperature, MinTemperature: currentMinimumTemperature, MaxTemperature: currentMaximumTemperature, Sunrise: currentSunrise, Sunset: currentSunset, UltraViolet: currentUV}
}

// main function to test all of the package
func main() {

	weatherObj := InitializeWeatherModule("Paris,Fr", "")

	fmt.Printf("Weather (" + weatherObj.Weather.GetMain() + ", " + weatherObj.Weather.GetDescription() + ", " + weatherObj.Weather.GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.Coords.GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.Coords.GetLatitude()) + ")\n")

	weatherObj.Temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + strconv.Itoa(weatherObj.UltraViolet.GetIndex()) + ", UV risk: " + weatherObj.UltraViolet.GetRisk().String() + "\n")

	weatherObj.Sunrise.SetCurrentFormatAsTimestamp()
	weatherObj.Sunset.SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + weatherObj.Sunrise.GetCurrentFormat().String() + "): " + weatherObj.Sunrise.GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + weatherObj.Sunset.GetCurrentFormat().String() + "): " + weatherObj.Sunset.GetSunTimeInCurrentFormat() + "\n")
}
