package weatherClasses

//
import (
	"fmt"
)

//
type PressureScale int

//
const (

	HectoPascal PressureScale = iota
	Pascal
	Bar
	Atmosphere
	Torr
	PoundsPerSquareInch
)

//
func (pressureUnitScale PressureScale) String() string {

	return [...]string{"hectoPascal", "pascal", "bar", "atmosphere", "torr", "poundsPerSquareInch"}[pressureUnitScale]
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
func (currentPressure *Pressure) SetPressureAsHectoPascal() {

	if currentPressure.scaleUnit == Pascal {

		currentPressure.value = currentPressure.value / 100
		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Bar {

		currentPressure.value = currentPressure.value * 1000
		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.value = currentPressure.value * 1013.2501
		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Torr {

		currentPressure.value = currentPressure.value * 1.333223684211
		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.value = currentPressure.value * 68.9475729318
		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else {

		fmt.Print(Red() + "Pressure already in HectoPascal (hPa)." + Reset() + "\n")

	}
}

// set pressure to Pascal (Pa)
func (currentPressure *Pressure) SetPressureAsPascal() {

	if currentPressure.scaleUnit == HectoPascal {

		currentPressure.value = currentPressure.value * 100
                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

		currentPressure.value = currentPressure.value * 100000
                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.value = currentPressure.value * 101325
                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

		currentPressure.value = currentPressure.value * 133.3223684211
                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.value = currentPressure.value * 6894.7572931783
                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Pascal (Pa)." + Reset() + "\n")

        }
}

// set pressure to Bar
func (currentPressure *Pressure) SetPressureAsBar() {

	if currentPressure.scaleUnit == HectoPascal {

		currentPressure.value = currentPressure.value / 1000
                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

		currentPressure.value = currentPressure.value / 100000
                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.value = currentPressure.value * 1.01325
                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

		currentPressure.value = currentPressure.value / 750.06375541921
                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.value = currentPressure.value / 14.5037738
                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Bar (bar)." + Reset() + "\n")

        }
}

// set pressure to Atmosphere
func (currentPressure *Pressure) SetPressureAsAtmosphere() {

	if currentPressure.scaleUnit == HectoPascal {

		currentPressure.value = currentPressure.value / 1013.25
                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

		currentPressure.value = currentPressure.value / 101325
                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

		currentPressure.value = currentPressure.value / 1.01325
                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

		currentPressure.value = currentPressure.value / 760
                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.value = currentPressure.value / 14.696
                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Atmosphere (atm)." + Reset() + "\n")

        }
}

// set pressure to Torr
func (currentPressure *Pressure) SetPressureAsTorr() {

	if currentPressure.scaleUnit == HectoPascal {

		currentPressure.value = currentPressure.value / 1.3332236842
                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

		currentPressure.value = currentPressure.value * 750.061682704
                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.value = currentPressure.value * 760
                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

		currentPressure.value = currentPressure.value / 133.3223684211
                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.value = currentPressure.value * 51.715
                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Torr (torr)." + Reset() + "\n")

        }
}

// set pressure to Pounds per square inch (psi)
func (currentPressure *Pressure) SetPressureAsPoundsPerSquareInch() {

	if currentPressure.scaleUnit == HectoPascal {

		currentPressure.value = currentPressure.value / 68.94757293168
                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

		currentPressure.value = currentPressure.value * 14.503773773022
                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.value = currentPressure.value * 14.695964
                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

		currentPressure.value = currentPressure.value / 6894.757293168
                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

		currentPressure.value = currentPressure.value / 51.715
                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Pounds per square inch (psi)." + Reset() + "\n")

        }
}

// 'value' attribute getter
func (currentPressure *Pressure) GetPressureValue() float64 {

	return currentPressure.value
}

// 'scaleUnit' attribute getter
func (currentPressure *Pressure) GetPressureScale() PressureScale {

	return currentPressure.scaleUnit
}

// 'symbolUnit' attribute getter
func (currentPressure *Pressure) GetPressureSymbolUnit() string {

	return currentPressure.symbolUnit
}
