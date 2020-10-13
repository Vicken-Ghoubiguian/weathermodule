package weatherClasses

//
type PressureScale int

//
const (

	HectoPascal PressureScale = iota
	Pascal
	Bar
	Atmosphere
	Torr
)

//
func (pressureUnitScale PressureScale) String() string {

	return [...]string{"hectoPascal", "pascal", "bar", "atmosphere", "torr"}[pressureUnitScale]
}

//
type Pressure struct {

	value float64
	scaleUnit PressureScale
	symbolUnit string
}

// pressure initializator with all needed parameters
func (currentPressure *Pressure) InitializePressure(pressureValue float64) {

	currentPressure.value = pressureValue
	currentPressure.scaleUnit = HectoPascal
	currentPressure.symbolUnit = "hPa"
}

// set pressure to HectoPascal (hPa)
func (currentPressure *Pressure) setPressureAsHectoPascal() {

}

// set pressure to Pascal (Pa)
func (currentPressure *Pressure) setPressureAsPascal() {

}

// set pressure to Bar
func (currentPressure *Pressure) setPressureAsBar() {

}

// set pressure to Atmosphere
func (currentPressure *Pressure) setPressureAsAtmosphere() {

}

// set pressure to Torr
func (currentPressure *Pressure) setPressureAsTorr() {

}

// set pressure to Pounds per square inch (psi)
func (currentPressure *Pressure) setPressureAsPoundsPerSquareInch() {

}

// 'value' attribute getter
func (currentPressure *Pressure) getPressureValue() float64 {

	return currentPressure.value
}

// 'scaleUnit' attribute getter
func (currentPressure *Pressure) getPressureScale() PressureScale {

	return currentPressure.scaleUnit
}

// 'symbolUnit' attribute getter
func (currentPressure *Pressure) getPressureSymbolUnit() string {

	return currentPressure.symbolUnit
}
