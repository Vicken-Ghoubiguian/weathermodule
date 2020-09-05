package weatherClasses

// Defining the type 'Coordinates' which stock and manage longitude and latitude of wished city
type Coordinates struct {

	//
	longitude float64
	latitude float64
}

// Defining the Coordinates initializer
func (coords *Coordinates) InitializeCoordinates(wishedongitude float64, wishedLatitude float64) {

	coords.longitude = wishedongitude
	coords.latitude = wishedLatitude
}

//
func (coords *Coordinates) GetLongitude() float64 {

	return coords.longitude
}

//
func (coords *Coordinates) GetLatitude() float64 {

	return coords.latitude
}
