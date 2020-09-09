package main

// import all necessary packages
import (
	weathermodule "../.."
	"fmt"
	"flag"
)

// main function to test all of the package
func main() {

	//
	var weatherObj weathermodule.WeatherModule

	//
	city := flag.String("city", "", "The city whose you want weather")
	countryCode := flag.String("countryCode", "", "The ISO 3166-1 alpha-2 or the ISO 3166-1 alpha-3 code of the wished country")
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")

	flag.Parse()

	weatherObj.InitializeWeatherModule(*city, *countryCode, *apiKey)

	fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")

	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.GetUltraViolet().GetIndex()) + ", UV risk: " + weatherObj.GetUltraViolet().GetRisk().String() + "\n")

	fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.GetWind().GetSpeed()) + "\n")
	fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.GetWind().GetDeg()) + "\n")
	fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.GetWind().GetGust()) + "\n")

	weatherObj.GetSunrise().SetCurrentFormatAsTimestamp()
	weatherObj.GetSunset().SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
}
