package weatherClasses

//
type Coordinates struct {

	//
	Longitude float64
	Latitude float64
}

//
func InitializeCoordinates(wishedongitude float64, wishedLatitude float64) *Coordinates {

	return &Coordinates{Longitude: wishedongitude, Latitude: wishedLatitude}
}
