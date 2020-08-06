package weatherClasses

//
type Coordinates struct {

	//
	longitude float64
	latitude float64
}

//
func InitializeCoordinates(wishedongitude float64, wishedLatitude float64) *Coordinates {

	return &Coordinates{longitude: wishedongitude, latitude: wishedLatitude}
}
