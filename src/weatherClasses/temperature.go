package weatherClasses

//
type TemperatureScaleEnum int

//
const(

	Kelvin TemperatureScaleEnum = 0
	Celsius TemperatureScaleEnum = 1
	Fahrenheit TemperatureScaleEnum = 2
)

//
type Temperature struct {

	//
	TemperatureValue float64
	TemperatureScale TemperatureScaleEnum
	TemperatureScaleSymbol string
}

//
func InitializeTemperature(value float64) *Temperature {

	return &Temperature{TemperatureValue: value, TemperatureScale: Kelvin, TemperatureScaleSymbol: " K"}
}
