package weatherClasses

//
type Weather struct {

	id int
	main string
	description string
	iconUrl string
}

//
func InitializeWeather(idValue int, mainValue string, descriptionValue string, icon string) *Weather {

	return &Weather{id: idValue, main: mainValue, description: descriptionValue, iconUrl: "https://openweathermap.org/img/wn/" + icon + ".png"}
}

//
func (currentWeather *Weather) GetId() int {

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
