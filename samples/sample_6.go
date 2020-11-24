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
	timeFormat := flag.String("timeFormat", "", "The datetime format to display sunrise and sunset time")
	dateSeparator := flag.String("dateSeparator", "\\", "The dateSeparator to display sunrise and sunset time")

	flag.Parse()

	weatherObj.InitializeMinimallyWeatherModule(*city, *apiKey)

	fmt.Printf("" + weatherObj.GetGeographicLocation().GetCityName() + " (" + weatherObj.GetGeographicLocation().GetCountryCode() + ")\n")

	fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")

	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.GetUltraViolet().GetIndex()) + ", UV risk: " + weatherObj.GetUltraViolet().GetRisk().String() + "\n")
	
	fmt.Printf("Humidity: " + fmt.Sprintf("%d", weatherObj.GetHumidity().GetHumidityValue()) + " " + weatherObj.GetHumidity().GetHumidityUnitScale() + "\n")

	fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.GetWind().GetSpeed()) + "\n")
	fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.GetWind().GetDeg()) + "\n")
	fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.GetWind().GetGust()) + "\n")

	if *timeFormat == "DMYHMS" {
	
		weatherObj.GetSunrise().SetCurrentDateSeparator(*dateSeparator)
		weatherObj.GetSunset().SetCurrentDateSeparator(*dateSeparator)

		weatherObj.GetSunrise().SetCurrentFormatAsDMYHMS()
		weatherObj.GetSunset().SetCurrentFormatAsDMYHMS()

		fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
        	fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
	
	} else if *timeFormat == "YMDHMS" {
	
		weatherObj.GetSunrise().SetCurrentDateSeparator(*dateSeparator)
		weatherObj.GetSunset().SetCurrentDateSeparator(*dateSeparator)

		weatherObj.GetSunrise().SetCurrentFormatAsYMDHMS()
		weatherObj.GetSunset().SetCurrentFormatAsYMDHMS()

		fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
		fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
	
	} else if *timeFormat == "MDYHMS" {
	
		weatherObj.GetSunrise().SetCurrentDateSeparator(*dateSeparator)
		weatherObj.GetSunset().SetCurrentDateSeparator(*dateSeparator)

		weatherObj.GetSunrise().SetCurrentFormatAsMDYHMS()
		weatherObj.GetSunset().SetCurrentFormatAsMDYHMS()

		fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
		fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
	
	} else if *timeFormat == "Timestamp" {
	
		weatherObj.GetSunrise().SetCurrentFormatAsTimestamp()
		weatherObj.GetSunset().SetCurrentFormatAsTimestamp()

		fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
		fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
	
	} else {

		fmt.Printf("\n" + "\033[31m" + "Error: the date and time format you entered does not exist..." + "\033[0m" + "\n\n")
	}
}
