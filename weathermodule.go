package weathermodule

// Import all necessary packages
import (
	"fmt"
	"strconv"
	"encoding/json"
	"usefullFunctions"
	"weatherClasses"
	"owmStructures"
	"io/ioutil"
	"net/http"
)

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
	humidity weatherClasses.Humidity

	//
	wind weatherClasses.Wind

	//
	sunrise weatherClasses.SunTime
	sunset weatherClasses.SunTime

	//
	geographicLocation weatherClasses.GeographicLocation

	//
	ultraViolet weatherClasses.UV
}

// Defining the Weather initializer
func (w *WeatherModule) InitializeWeatherModule(city string, countrysISOAlpha2Code string, apiKey string) {

	//
	var owm owmStructures.OWMStruct

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city + "," + countrysISOAlpha2Code, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//Single instruction to convert weather_json_string []byte variable to string
	err = json.Unmarshal(weatherJsonString, &owm)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	if owm.Cod != 200 {

		// Calling the 'owmErrorHandlerFunction' from the 'usefullFunctions' module to treat the current error...
		usefullFunctions.OwmErrorHandlerFunction(strconv.Itoa(int(owm.Cod)), weatherClasses.Red(), owm.Message, weatherClasses.Reset())

	} else {

		//
		var UVowm owmStructures.UVStruct

		//
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, fmt.Sprintf("%g", owm.Coord.Lat), fmt.Sprintf("%g", owm.Coord.Lon))

		//
		uvResp, err := http.Get(uvRequest)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		//weather := usefullFunctions.ExtractWeatherFromJSONFunction(gjson.Get(string(weatherJsonString), "weather").String())
		err = json.Unmarshal(uvJsonString, &UVowm)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		w.coords.InitializeCoordinates(owm.Coord.Lon, owm.Coord.Lat)

		//
		w.weather.InitializeWeather(owm.Weather[0].Id, owm.Weather[0].Main, owm.Weather[0].Description, owm.Weather[0].Icon)

		//
		w.temperature.InitializeTemperature(owm.Main.Temp)
		w.feelingLikeTemperature.InitializeTemperature(owm.Main.Feels_like)
		w.minTemperature.InitializeTemperature(owm.Main.Temp_min)
		w.maxTemperature.InitializeTemperature(owm.Main.Temp_max)

		//
		w.wind.InitializeWind(owm.Wind.Speed, owm.Wind.Deg, owm.Wind.Gust)

		//
		w.pressure.InitializePressure(owm.Main.Pressure)

		//
		w.humidity.InitializeHumidity(owm.Main.Humidity)

		//
		w.geographicLocation.InitializeGeographicLocation(owm.Sys.Country, owm.Name)

		//
		w.sunrise.InitializeSunTime(owm.Sys.Sunrise)
		w.sunset.InitializeSunTime(owm.Sys.Sunset)

		//
		w.ultraViolet.InitializeUV(int64(UVowm.Value))

		// Displaying success message...
		fmt.Println(weatherClasses.Green() + "Weather implemented successfully !" + weatherClasses.Reset() + "\n")
	}
}

// Defining the minimal Weather initializer
func (w *WeatherModule) InitializeMinimallyWeatherModule(city string, apiKey string) {

	//
	var owm owmStructures.OWMStruct

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//Single instruction to convert weather_json_string []byte variable to string
	err = json.Unmarshal(weatherJsonString, &owm)

	//
	usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

	//
	if owm.Cod != 200 {

		// Calling the 'owmErrorHandlerFunction' from the 'usefullFunctions' module to treat the current error...
		usefullFunctions.OwmErrorHandlerFunction(strconv.Itoa(int(owm.Cod)), weatherClasses.Red(), owm.Message, weatherClasses.Reset())

	} else {

		//
		var UVowm owmStructures.UVStruct

		//
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, fmt.Sprintf("%g", owm.Coord.Lat), fmt.Sprintf("%g", owm.Coord.Lon))

		//
		uvResp, err := http.Get(uvRequest)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		err = json.Unmarshal(uvJsonString, &UVowm)

		//
		usefullFunctions.OtherErrorHandlerFunction(err, weatherClasses.Red(), weatherClasses.Reset())

		//
		w.coords.InitializeCoordinates(owm.Coord.Lon, owm.Coord.Lat)

		//
		w.weather.InitializeWeather(owm.Weather[0].Id, owm.Weather[0].Main, owm.Weather[0].Description, owm.Weather[0].Icon)

		//
		w.temperature.InitializeTemperature(owm.Main.Temp)
		w.feelingLikeTemperature.InitializeTemperature(owm.Main.Feels_like)
		w.minTemperature.InitializeTemperature(owm.Main.Temp_min)
		w.maxTemperature.InitializeTemperature(owm.Main.Temp_max)

		//
		w.wind.InitializeWind(owm.Wind.Speed, owm.Wind.Deg, owm.Wind.Gust)

		//
		w.pressure.InitializePressure(owm.Main.Pressure)

		//
		w.humidity.InitializeHumidity(owm.Main.Humidity)

		//
		w.geographicLocation.InitializeGeographicLocation(owm.Sys.Country, owm.Name)

		//
		w.sunrise.InitializeSunTime(owm.Sys.Sunrise)
		w.sunset.InitializeSunTime(owm.Sys.Sunset)

		//
		w.ultraViolet.InitializeUV(int64(UVowm.Value))

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
func (w *WeatherModule) GetHumidity() *weatherClasses.Humidity {

	return &w.humidity
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
func (w *WeatherModule) GetGeographicLocation() *weatherClasses.GeographicLocation {

	return &w.geographicLocation
}

//
func (w *WeatherModule) GetUltraViolet() *weatherClasses.UV {

	return &w.ultraViolet
}
