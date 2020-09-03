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
	city := flag.String("city", "", "The city whose you want weather")
	countryCode := flag.String("countryCode", "", "The ISO 3166-1 alpha-2 or the ISO 3166-1 alpha-3 code of the wished country")
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")

	flag.Parse()

	weatherObj := weathermodule.InitializeWeatherModule(*city, *countryCode, *apiKey)

	fmt.Printf("Weather (" + weatherObj.Weather.GetMain() + ", " + weatherObj.Weather.GetDescription() + ", " + weatherObj.Weather.GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.Coords.GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.Coords.GetLatitude()) + ")\n")

	weatherObj.Temperature.SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	weatherObj.Temperature.SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.Temperature.GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.Temperature.GetTemperatureValue()) + weatherObj.Temperature.GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.UltraViolet.GetIndex()) + ", UV risk: " + weatherObj.UltraViolet.GetRisk().String() + "\n")

	fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.Wind.GetSpeed()) + "\n")
	fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.Wind.GetDeg()) + "\n")
	fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.Wind.GetGust()) + "\n")

	weatherObj.Sunrise.SetCurrentFormatAsTimestamp()
	weatherObj.Sunset.SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + weatherObj.Sunrise.GetCurrentFormat().String() + "): " + weatherObj.Sunrise.GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + weatherObj.Sunset.GetCurrentFormat().String() + "): " + weatherObj.Sunset.GetSunTimeInCurrentFormat() + "\n")
}
