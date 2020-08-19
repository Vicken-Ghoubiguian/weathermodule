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
	/*owmCode := gjson.Get(string(weatherJsonString), "cod")

	//
	if owmCode.Int() != 200 {

		//
		owmMessage := gjson.Get(string(weatherJsonString), "message")

		// Calling the 'owmErrorHandlerFunction' to treat the current error...
		owmErrorHandlerFunction(owmCode.String(), owmMessage.String())

	} else {*/

		/*******************************************************************

                ********************************************************************/

		//
		longitude := gjson.Get(string(weatherJsonString), "coord.lon")
		latitude := gjson.Get(string(weatherJsonString), "coord.lat")

		//
		weather := extractWeatherFromJSONFunction(gjson.Get(string(weatherJsonString), "weather").String())
		id := gjson.Get(weather, "id")
		main := gjson.Get(weather, "main")
		description := gjson.Get(weather, "description")
		icon := gjson.Get(weather, "icon")

		//
		temperature := gjson.Get(string(weatherJsonString), "main.temp")
		feelingTemperature := gjson.Get(string(weatherJsonString), "main.feels_like")
		minimumTemperature := gjson.Get(string(weatherJsonString), "main.temp_min")
		maximumTemperature := gjson.Get(string(weatherJsonString), "main.temp_max")

		//
		sunrise := gjson.Get(string(weatherJsonString), "sys.sunrise")
		sunset := gjson.Get(string(weatherJsonString), "sys.sunset")

		/*******************************************************************

		********************************************************************/

		//
		currentCoordinates := weatherClasses.InitializeCoordinates(longitude.Float(), latitude.Float())

		//
		currentWeather := weatherClasses.InitializeWeather(id.Int(), main.String(), description.String(), icon.String())

		//
		currentTemperature := weatherClasses.InitializeTemperature(temperature.Float())
		currentFeelingTemperature := weatherClasses.InitializeTemperature(feelingTemperature.Float())
		currentMinimumTemperature := weatherClasses.InitializeTemperature(minimumTemperature.Float())
		currentMaximumTemperature := weatherClasses.InitializeTemperature(maximumTemperature.Float())

		//
		currentSunrise := weatherClasses.InitializeSunTime(sunrise.Int())
		currentSunset := weatherClasses.InitializeSunTime(sunset.Int())

		//
		currentUV := weatherClasses.InitializeUV(10)

		// Displaying success message...
		fmt.Println(green + "Weather implemented successfully !" + reset + "\n")

		return &WeatherModule{Coords: currentCoordinates, Weather: currentWeather, Temperature: currentTemperature, FeelingLikeTemperature: currentFeelingTemperature, MinTemperature: currentMinimumTemperature, MaxTemperature: currentMaximumTemperature, Sunrise: currentSunrise, Sunset: currentSunset, UltraViolet: currentUV}
	//}

	//return string(weatherJsonString), string(uvJsonString)
}

// main function to test all of the package
func main() {

	coords := weatherClasses.InitializeCoordinates(3.45, 7.89)
	temperature := weatherClasses.InitializeTemperature(300.85)
	uv := weatherClasses.InitializeUV(10)
	sunRise := weatherClasses.InitializeSunTime(1597221424)
	sunSet := weatherClasses.InitializeSunTime(1597221424)
	weatherObj := InitializeWeatherModule("Paris,Fr", "")

	/*fmt.Printf("Weather response: " + weatherResponse + "\n\n")
	fmt.Printf("UV response: " + uvResponse + "\n\n")*/

	fmt.Printf("Weather (" + weatherObj.Weather.GetMain() + ", " + weatherObj.Weather.GetDescription() + ", " + weatherObj.Weather.GetIconUrl() + ")\n")

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
