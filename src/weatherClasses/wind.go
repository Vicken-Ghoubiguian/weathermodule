package weatherClasses

//
type Wind struct {

	windSpeed float64
	windDeg int64
	windGust float64
}

// Defining the Wind initializer
func InitializeWind(wishedSpeed float64, wishedDeg int64, wishedGust float64) *Wind {

	return &Wind{windSpeed: wishedSpeed, windDeg: wishedDeg, windGust: wishedGust}
}

//
func (currentWind *Wind) GetSpeed() float64 {

	return currentWind.windSpeed
}

//
func (currentWind *Wind) GetDeg() int64 {

	return currentWind.windDeg
}

//
func (currentWind *Wind) GetGust() float64 {

	return currentWind.windGust
}
