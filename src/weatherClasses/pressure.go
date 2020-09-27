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

//
func (currentPressure *Pressure) InitializePressure(pressureValue float64) {

	currentPressure.value = pressureValue
	currentPressure.scaleUnit = HectoPascal
	currentPressure.symbolUnit = "hPa"
}

//
func (currentPressure *Pressure) getPressureValue() float64 {

	return currentPressure.value
}

//
func (currentPressure *Pressure) getPressureScale() PressureScale {

	return currentPressure.scaleUnit
}

//
func (currentPressure *Pressure) getPressureSymbolUnit() string {

	return currentPressure.symbolUnit
}
