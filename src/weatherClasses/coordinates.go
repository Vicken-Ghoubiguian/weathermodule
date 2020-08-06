package weatherClasses

// Defining the type 'Coordinates' which stock and manage longitude and latitude of wished city
type Coordinates struct {

	//
	Longitude float64
	Latitude float64
}

// Defining the Coordinates initializer
func InitializeCoordinates(wishedongitude float64, wishedLatitude float64) *Coordinates {

	return &Coordinates{Longitude: wishedongitude, Latitude: wishedLatitude}
}
