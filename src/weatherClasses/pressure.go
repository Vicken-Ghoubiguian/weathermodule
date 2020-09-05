package weatherClasses

//
type Pressure struct {

	value float64
	unit string
}

//
func (currentPressure *Pressure) InitializePressure(pressureValue float64) {

	currentPressure.value = pressureValue
	currentPressure.unit = "hPa"
}

//
func (currentPressure *Pressure) InitializePressureWithUnit(pressureValue float64, pressureUnit string) {

	currentPressure.value = pressureValue
	currentPressure.unit = pressureUnit
}
