package weathermodule

// Import all necessary packages
import (
	"fmt"
	"usefulFunctions"
	"weatherClasses"
	"io/ioutil"
	"net/http"
)

// Importation of the github project gjson to treat received json
import "github.com/tidwall/gjson"

// Defining the type 'Weather' which recover and manage current weather in a defined city
type WeatherModule struct {

	//
	coords weatherClasses.Coordinates

	//
	weather weatherClasses.Weather

	//
	temperature weatherClasses.Temperature
	feelingLikeTemperature weatherClasses.Temperature
	minTemperature weatherClasses.Temperature
	maxTemperature weatherClasses.Temperature

	//
	pressure weatherClasses.Pressure

	//
	humidity int64

	//
	wind weatherClasses.Wind

	//
	sunrise weatherClasses.SunTime
	sunset weatherClasses.SunTime

	//
	countryCode string
	city string

	//
	currentLanguage weatherClasses.Language

	//
	ultraViolet weatherClasses.UV
}

// Defining the Weather initializer
func (w *WeatherModule) InitializeWeatherModule(city string, countrysISOAlpha2Code string, apiKey string) {

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city + "," + countrysISOAlpha2Code, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	usefulFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	usefulFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	owmCode := gjson.Get(string(weatherJsonString), "cod")

	//
	if owmCode.Int() != 200 {

		//
		owmMessage := gjson.Get(string(weatherJsonString), "message")

		// Calling the 'owmErrorHandlerFunction' from the 'usefulFunctions' module to treat the current error...
		usefulFunctions.OwmErrorHandlerFunction(owmCode.String(), weatherClasses.Red(), owmMessage.String(), weatherClasses.Reset())

	} else {

		//
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, gjson.Get(string(weatherJsonString), "coord.lat").String(), gjson.Get(string(weatherJsonString), "coord.lon").String())

		//
		uvResp, err := http.Get(uvRequest)

		//
		usefulFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		usefulFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		weather := usefulFunctions.ExtractWeatherFromJSONFunction(gjson.Get(string(weatherJsonString), "weather").String())

		//
		w.coords.InitializeCoordinates(gjson.Get(string(weatherJsonString), "coord.lon").Float(), gjson.Get(string(weatherJsonString), "coord.lat").Float())

		//
		w.weather.InitializeWeather(gjson.Get(weather, "id").Int(), gjson.Get(weather, "main").String(), gjson.Get(weather, "description").String(), gjson.Get(weather, "icon").String())

		//
		w.temperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp").Float())
		w.feelingLikeTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.feels_like").Float())
		w.minTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_min").Float())
		w.maxTemperature.InitializeTemperature(gjson.Get(string(weatherJsonString), "main.temp_max").Float())

		//
		w.wind.InitializeWind(gjson.Get(string(weatherJsonString), "wind.speed").Float(), gjson.Get(string(weatherJsonString), "wind.deg").Int(), gjson.Get(string(weatherJsonString), "wind.gust").Float())

		//
		w.pressure.InitializePressure(gjson.Get(string(weatherJsonString), "main.pressure").Float())

		//
		w.humidity = gjson.Get(string(weatherJsonString), "main.humidity").Int()

		//
		w.countryCode = gjson.Get(string(weatherJsonString), "sys.country").String()

		//
		w.city = gjson.Get(string(weatherJsonString), "name").String()

		//
		w.sunrise.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunrise").Int())
		w.sunset.InitializeSunTime(gjson.Get(string(weatherJsonString), "sys.sunset").Int())

		//
		w.ultraViolet.InitializeUV(gjson.Get(string(uvJsonString), "value").Int())

		// Displaying success message...
		fmt.Println(weatherClasses.Green() + "Weather implemented successfully !" + weatherClasses.Reset() + "\n")
	}
}

//
func (w *WeatherModule) GetCoords() *weatherClasses.Coordinates {

	return &w.coords
}

//
func (w *WeatherModule) GetWeather() *weatherClasses.Weather {

	return &w.weather
}

//
func (w *WeatherModule) GetTemperature() *weatherClasses.Temperature {

	return &w.temperature
}

//
func (w *WeatherModule) GetFeelingLikeTemperature() *weatherClasses.Temperature {

	return &w.feelingLikeTemperature
}

//
func (w *WeatherModule) GetMinTemperature() *weatherClasses.Temperature {

	return &w.minTemperature
}

//
func (w *WeatherModule) GetMaxTemperature() *weatherClasses.Temperature {

	return &w.maxTemperature
}

//
func (w *WeatherModule) GetPressure() *weatherClasses.Pressure {

	return &w.pressure
}

//
func (w *WeatherModule) GetHumidity() int64 {

	return w.humidity
}

//
func (w *WeatherModule) GetWind() *weatherClasses.Wind {

	return &w.wind
}

//
func (w *WeatherModule) GetSunrise() *weatherClasses.SunTime {

	return &w.sunrise
}

//
func (w *WeatherModule) GetSunset() *weatherClasses.SunTime {

	return &w.sunset
}

//
func (w *WeatherModule) GetCountryCode() string {

	return w.countryCode
}

//
func (w *WeatherModule) GetCity() string {

	return w.city
}

//
func (w *WeatherModule) GetLanguage() *weatherClasses.Language {

	return &w.currentLanguage
}

//
func (w *WeatherModule) GetUltraViolet() *weatherClasses.UV {

	return &w.ultraViolet
}
