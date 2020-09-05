package weatherClasses

//
type Pressure struct {

	value float64
	unit string
}

//
func InitializePressure(pressureValue float64) *Pressure {

	return &Pressure{value: pressureValue, unit: "hPa"}
}

//
func InitializePressureWithUnit(pressureValue float64, pressureUnit string) *Pressure {

	return &Pressure{value: pressureValue, unit: pressureUnit}
}
