package weatherClasses

//
type TemperatureScale int

//
const (

	Kelvin TemperatureScale = iota
	Celsius
	Fahrenheit
)

//
func (tempScale TemperatureScale) String() string {

	return [...]string{"Kelvin", "Celsius", "Fahrenheit"}[tempScale]
}

//
type Temperature struct {

	//
	TemperatureValue float64
	CurrentTemperatureScale TemperatureScale
	TemperatureScaleSymbol string
}

//
func InitializeTemperature(value float64) *Temperature {

	return &Temperature{TemperatureValue: value, CurrentTemperatureScale: Kelvin, TemperatureScaleSymbol: " K"}
}
