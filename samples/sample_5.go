package main

// import all necessary packages
import (
	weathermodule ".."
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
	temperatureScaleValue := *temperatureScale
	
	//
	weatherObj.InitializeWeatherModule(*city, *countryCode, *apiKey)

	//
	fmt.Printf("" + weatherObj.GetGeographicLocation().GetCityName() + " (" + weatherObj.GetGeographicLocation().GetCountryCode() + ")\n")

	//
	fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

	//
	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")

	if temperatureScaleValue == "Celsius" || temperatureScaleValue == "Fahrenheit" || temperatureScaleValue == "Kelvin"{

		fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		if temperatureScaleValue == "Celsius" {
		
			weatherObj.GetTemperature().SetTemperatureAsCelsius()
			weatherObj.GetFeelingLikeTemperature().SetTemperatureAsCelsius()
			weatherObj.GetMinTemperature().SetTemperatureAsCelsius()
			weatherObj.GetMaxTemperature().SetTemperatureAsCelsius()
		
		} else if temperatureScaleValue == "Fahrenheit" {
		
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

		fmt.Printf("Feeling temperature (in " + weatherObj.GetFeelingLikeTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetFeelingLikeTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

		fmt.Printf("Min temperature (in " + weatherObj.GetMinTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetMinTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")
		
		fmt.Printf("Max temperature (in " + weatherObj.GetMaxTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetMaxTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	} else {

		fmt.Printf("\n" + "\033[31m" + "Error: the temperature scale you entered does not exist..." + "\033[0m" + "\n\n")
	}

	fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.GetUltraViolet().GetIndex()) + ", UV risk: " + weatherObj.GetUltraViolet().GetRisk().String() + "\n")

	fmt.Printf("Humidity: " + fmt.Sprintf("%d", weatherObj.GetHumidity().GetHumidityValue()) + " " + weatherObj.GetHumidity().GetHumidityUnitScale() + "\n")

	fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.GetWind().GetSpeed()) + "\n")
	fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.GetWind().GetDeg()) + "\n")
	fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.GetWind().GetGust()) + "\n")

	weatherObj.GetSunrise().SetCurrentFormatAsTimestamp()
	weatherObj.GetSunset().SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
}
