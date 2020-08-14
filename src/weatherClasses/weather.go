package weatherClasses

//
type Weather struct {

	id int
	main string
	description string
	iconUrl string
}

//
func InitializeWeather(idValue int, mainValue string, descriptionValue string, icon string) {

	return &Weather{id: idValue, main: mainValue, description: descriptionValue, iconUrl: "https://openweathermap.org/img/wn/" + icon + ".png"}
}
