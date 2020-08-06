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
	temperatureValue float64
	currentTemperatureScale TemperatureScale
	temperatureScaleSymbol string
}

// Defining the Temperature initializer
func InitializeTemperature(value float64) *Temperature {

	return &Temperature{temperatureValue: value, currentTemperatureScale: Kelvin, temperatureScaleSymbol: " K"}
}

// 'temperatureValue' attribute getter
func (temp *Temperature) GetTemperatureValue() float64 {

	return temp.temperatureValue
}

// 'currentTemperatureScale' attribute getter
func (temp *Temperature) GetCurrentTemperatureScale() TemperatureScale {

	return temp.currentTemperatureScale
}

// 'temperatureScaleSymbol' attribute getter
func (temp *Temperature) GetTemperatureScaleSymbol() string {

	return temp.temperatureScaleSymbol
}
