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
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")
	pressureScale := flag.String("pressureScale", "", "The choosen pressure scale")

	flag.Parse()

	weatherObj.InitializeMinimallyWeatherModule(*city, *apiKey)

	fmt.Printf("" + weatherObj.GetGeographicLocation().GetCityName() + " (" + weatherObj.GetGeographicLocation().GetCountryCode() + ")\n")

	fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")
	
	fmt.Printf("\n")

	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")
	
	fmt.Printf("\n")
	
	//
	pressureScaleValue := *pressureScale
	
	if pressureScaleValue == "HectoPascal" || pressureScaleValue == "Pascal" || pressureScaleValue == "Bar" || pressureScaleValue == "Atmosphere" || pressureScaleValue == "Torr" || pressureScaleValue == "PoundsPerSquareInch"{
	
		if pressureScaleValue == "Pascal" {
		
			weatherObj.GetPressure().SetPressureAsPascal()
		
		} else if pressureScaleValue == "Bar" {
		
			weatherObj.GetPressure().SetPressureAsBar()
			
		} else if pressureScaleValue == "Atmosphere" {
		
			weatherObj.GetPressure().SetPressureAsAtmosphere()
		
		} else if pressureScaleValue == "Torr" {
		
			weatherObj.GetPressure().SetPressureAsTorr()
		
		} else if pressureScaleValue == "PoundsPerSquareInch" {
		
			weatherObj.GetPressure().SetPressureAsPoundsPerSquareInch()
		
		} else if pressureScaleValue == "HectoPascal" {
		
			weatherObj.GetPressure().SetPressureAsHectoPascal()

		} else {

			fmt.Printf("\n" + "\033[31m" + "Error: the pressure scale you entered does not exist..." + "\033[0m" + "\n\n")
		}
	}

	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + " " + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	fmt.Printf("\n")

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
