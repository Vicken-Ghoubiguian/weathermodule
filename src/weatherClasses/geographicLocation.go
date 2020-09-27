package weatherClasses

//
type GeographicLocation struct {

	countryCode string
	cityName string
}

//
func (currentGeographicLocation *GeographicLocation) InitializeGeographicLocation(countryCode string, cityName string) {

	currentGeographicLocation.countryCode = countryCode
	currentGeographicLocation.cityName = cityName
}

//
func (currentGeographicLocation *GeographicLocation) GetCountryCode() string {

	return currentGeographicLocation.countryCode
}

//
func (currentGeographicLocation *GeographicLocation) GetCityName() string {

	return currentGeographicLocation.cityName
}
