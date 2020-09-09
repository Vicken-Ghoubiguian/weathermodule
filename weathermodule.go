package weathermodule

// Import all necessary packages
import (
	"fmt"
	"os"
	"weatherClasses"
	"io/ioutil"
	"net/http"
	"strings"
)

// Importation of the github project gjson to treat received json
import "github.com/tidwall/gjson"

// Function which extracts weather datas from JSON response
func extractWeatherFromJSONFunction(weatherFromHTTPResponseString string) string {

	brutWeatherWithoutHooks := strings.Trim(weatherFromHTTPResponseString, "[]")

	return brutWeatherWithoutHooks
}

// Function which display HTTP request error's code and message when the first occurs
func owmErrorHandlerFunction(codeError string, errorMessage string) {

	fmt.Println(weatherClasses.Red() + "Occured error (" + codeError + "): " + errorMessage + weatherClasses.Reset())

	fmt.Println("\n")

	os.Exit(1)
}

// Function which display other errors when they occurs
func otherErrorHandlerFunction(err error) {

	if err != nil {

		fmt.Println(weatherClasses.Red() + err.Error() + weatherClasses.Reset())

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
	Pressure weatherClasses.Pressure

	//
	Humidity int64

	//
	Wind weatherClasses.Wind

	//
	Sunrise weatherClasses.SunTime
	Sunset weatherClasses.SunTime

	//
	UltraViolet weatherClasses.UV
}

// Defining the Weather initializer
func (w *WeatherModule) InitializeWeatherModule(city string, countrysISOAlpha2Code string, apiKey string) {

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city + "," + countrysISOAlpha2Code, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	otherErrorHandlerFunction(err)

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	otherErrorHandlerFunction(err)

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
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, gjson.Get(string(weatherJsonString), "coord.lat").String(), gjson.Get(string(weatherJsonString), "coord.lon").String())

		//
		uvResp, err := http.Get(uvRequest)

		//
		otherErrorHandlerFunction(err)

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		otherErrorHandlerFunction(err)

		//
		weather := extractWeatherFromJSONFunction(gjson.Get(string(weatherJsonString), "weather").String())

		//
		w.Coords.InitializeCoordinates(gjson.Get(string(weatherJsonString), "coord.lon").Float(), gjson.Get(string(weatherJsonString), "coord.lat").Float())

		//
		w.Weather.InitializeWeather(gjson.Get(weather, "id").Int(), gjson.Get(weather, "main").String(), gjson.Get(weather, "description").String(), gjson.Get(weather, "icon").String())

		//
		w.Temperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp").Float())
		w.FeelingLikeTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.feels_like").Float())
		w.MinTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_min").Float())
		w.MaxTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_max").Float())

		//
		w.Wind.InitializeWind(gjson.Get(string(weatherJsonString), "wind.speed").Float(), gjson.Get(string(weatherJsonString), "wind.deg").Int(), gjson.Get(string(weatherJsonString), "wind.gust").Float())

		//
		w.Pressure.InitializePressure(gjson.Get(string(weatherJsonString), "main.pressure").Float())

		//
		w.Humidity = gjson.Get(string(weatherJsonString), "main.humidity").Int()

		//
		w.Sunrise.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunrise").Int())
		w.Sunset.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunset").Int())

		//
		w.UltraViolet.InitializeUV(gjson.Get(string(uvJsonString), "value").Int())

		// Displaying success message...
		fmt.Println(weatherClasses.Green() + "Weather implemented successfully !" + weatherClasses.Reset() + "\n")
	}
}

//
func (w *WeatherModule) GetCoords() weatherClasses.Coordinates {

	return w.Coords
}

//
func (w *WeatherModule) GetWeather() weatherClasses.Weather {

	return w.Weather
}
