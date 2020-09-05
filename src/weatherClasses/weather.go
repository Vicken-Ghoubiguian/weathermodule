package weatherClasses

//
type Weather struct {

	id int64
	main string
	description string
	iconUrl string
}

//
func (currentWeather *Weather) InitializeWeather(idValue int64, mainValue string, descriptionValue string, icon string) {

	//return &Weather{id: idValue, main: mainValue, description: descriptionValue, iconUrl: "https://openweathermap.org/img/wn/" + icon + ".png"}
	currentWeather.id = idValue
	currentWeather.main = mainValue
	currentWeather.description = descriptionValue
	currentWeather.iconUrl = "https://openweathermap.org/img/wn/" + icon + ".png"
}

//
func (currentWeather *Weather) GetId() int64 {

	return currentWeather.id
}

//
func (currentWeather *Weather) GetMain() string {

	return currentWeather.main
}

//
func (currentWeather *Weather) GetDescription() string {

	return currentWeather.description
}

//
func (currentWeather *Weather) GetIconUrl() string {

	return currentWeather.iconUrl
}
