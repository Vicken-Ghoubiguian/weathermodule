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
	temperatureScale := flag.String("temperatureScale", "", "The choosen temperature scale")

	flag.Parse()
	
	//
	weatherObj.InitializeWeatherModule(*city, *countryCode, *apiKey)
	
	if temperatureScale == "Celsius" || temperatureScale == "Fahrenheit" || temperatureScale == "Kelvin"{

		fmt.Printf("" + weatherObj.GetCity() + " (" + weatherObj.GetCountryCode() + ")\n")

		fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

		fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")

		fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		if temperatureScale == "Celsius" {
		
			weatherObj.GetTemperature().SetTemperatureAsCelsius()
			weatherObj.GetFeelingLikeTemperature().SetTemperatureAsCelsius()
			weatherObj.GetMinTemperature().SetTemperatureAsCelsius()
			weatherObj.GetMaxTemperature().SetTemperatureAsCelsius()
		
		} else if temperatureScale == "Fahrenheit" {
		
			weatherObj.GetTemperature().SetTemperatureAsFahrenheit()
			weatherObj.GetFeelingLikeTemperature().SetTemperatureAsFahrenheit()
			weatherObj.GetMinTemperature().SetTemperatureAsFahrenheit()
			weatherObj.GetMaxTemperature().SetTemperatureAsFahrenheit()
		
		} else {
		
			weatherObj.GetTemperature().SetTemperatureAsKelvin()
			weatherObj.GetFeelingLikeTemperature().SetTemperatureAsKelvin()
			weatherObj.GetMinTemperature().SetTemperatureAsKelvin()
			weatherObj.GetMaxTemperature().SetTemperatureAsKelvin()
		
		}


		fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		fmt.Printf("Temperature (in " + weatherObj.GetFeelingLikeTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetFeelingLikeTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		fmt.Printf("Temperature (in " + weatherObj.GetMinTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetMinTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")
		
		fmt.Printf("Temperature (in " + weatherObj.GetMaxTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetMaxTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.GetUltraViolet().GetIndex()) + ", UV risk: " + weatherObj.GetUltraViolet().GetRisk().String() + "\n")

		fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.GetWind().GetSpeed()) + "\n")
		fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.GetWind().GetDeg()) + "\n")
		fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.GetWind().GetGust()) + "\n")

		weatherObj.GetSunrise().SetCurrentFormatAsTimestamp()
		weatherObj.GetSunset().SetCurrentFormatAsTimestamp()

		fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
		fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
	} 
}
