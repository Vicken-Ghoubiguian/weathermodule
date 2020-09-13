package usefulFunctions

import(
	"strings"
)

// Function which extracts weather datas from JSON response
func ExtractWeatherFromJSONFunction(weatherFromHTTPResponseString string) string {

        brutWeatherWithoutHooks := strings.Trim(weatherFromHTTPResponseString, "[]")

        return brutWeatherWithoutHooks
}
