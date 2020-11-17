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
func (currentPressure *Pressure) setPressureAsHectoPascal() {

	if currentPressure.scaleUnit == Pascal {

		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Bar {

		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Atmosphere {

		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == Torr {

		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else if currentPressure.scaleUnit == PoundsPerSquareInch {

		currentPressure.scaleUnit = HectoPascal
		currentPressure.symbolUnit = " hPa"

		fmt.Print(Green() + "Pressure converted in HectoPascal (hPa) successfully." + Reset() + "\n")

	} else {

		fmt.Print(Red() + "Pressure already in HectoPascal (hPa)." + Reset() + "\n")

	}
}

// set pressure to Pascal (Pa)
func (currentPressure *Pressure) setPressureAsPascal() {

	if currentPressure.scaleUnit == HectoPascal {

                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

                currentPressure.scaleUnit = Pascal
                currentPressure.symbolUnit = " Pa"

                fmt.Print(Green() + "Pressure converted in Pascal (Pa) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Pascal (Pa)." + Reset() + "\n")

        }
}

// set pressure to Bar
func (currentPressure *Pressure) setPressureAsBar() {

	if currentPressure.scaleUnit == HectoPascal {

                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

                currentPressure.scaleUnit = Bar
                currentPressure.symbolUnit = " bar"

                fmt.Print(Green() + "Pressure converted in Bar (bar) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Bar (bar)." + Reset() + "\n")

        }
}

// set pressure to Atmosphere
func (currentPressure *Pressure) setPressureAsAtmosphere() {

	if currentPressure.scaleUnit == HectoPascal {

                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

                currentPressure.scaleUnit = Atmosphere
                currentPressure.symbolUnit = " atm"

                fmt.Print(Green() + "Pressure converted in Atmosphere (atm) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Atmosphere (atm)." + Reset() + "\n")

        }
}

// set pressure to Torr
func (currentPressure *Pressure) setPressureAsTorr() {

	if currentPressure.scaleUnit == HectoPascal {

                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == PoundsPerSquareInch {

                currentPressure.scaleUnit = Torr
                currentPressure.symbolUnit = " torr"

                fmt.Print(Green() + "Pressure converted in Torr (torr) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Torr (torr)." + Reset() + "\n")

        }
}

// set pressure to Pounds per square inch (psi)
func (currentPressure *Pressure) setPressureAsPoundsPerSquareInch() {

	if currentPressure.scaleUnit == HectoPascal {

                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Bar {

                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Atmosphere {

                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Pascal {

                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else if currentPressure.scaleUnit == Torr {

                currentPressure.scaleUnit = PoundsPerSquareInch
                currentPressure.symbolUnit = " psi"

                fmt.Print(Green() + "Pressure converted in Pounds per square inch (psi) successfully." + Reset() + "\n")

        } else {

                fmt.Print(Red() + "Pressure already in Pounds per square inch (psi)." + Reset() + "\n")

        }
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
