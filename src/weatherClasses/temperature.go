package weatherClasses

//
type Temperature struct {

	//
	TemperatureValue float64
	CurrentTemperatureScale string
	TemperatureScaleSymbol string
}

//
func InitializeTemperature(value float64) *Temperature {

	return &Temperature{TemperatureValue: value, CurrentTemperatureScale: "Kelvin", TemperatureScaleSymbol: " K"}
}
