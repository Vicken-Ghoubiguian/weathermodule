package weatherClasses

//
import (
	"fmt"
	"time"
	"strconv"
)

//
type sunTimeFormat int

//
const (

	Timestamp = iota
	DMYHMSFormat
	YMDHMSFormat
	MDYHMSFormat
)

//
func (timeFormat sunTimeFormat) String() string {

	return [...]string{"Timestamp", "DMYHMSFormat", "YMDHMSFormat", "MDYHMSFormat"}[timeFormat]
}

//
func timePresentationFunction(examinedInt int) string {

	if examinedInt < 10 {
	
		return "0" + strconv.Itoa(examinedInt)
		
	} else {
		
		return strconv.Itoa(examinedInt)
	}
}

//
type SunTime struct {

	//
	sunTimeInCurrentFormat string
	currentFormat sunTimeFormat
	currentDateSeparator string
	asTimestamp int64
}

//
func (currentSunTime *SunTime) InitializeSunTime(value int64) {

	currentSunTime.sunTimeInCurrentFormat = strconv.FormatInt(value, 10)
	currentSunTime.currentFormat = Timestamp
	currentSunTime.currentDateSeparator = ""
	currentSunTime.asTimestamp = value
}

//
func (currentSunTime *SunTime) GetSunTimeInCurrentFormat() string {

	return currentSunTime.sunTimeInCurrentFormat
}

//
func (currentSunTime *SunTime) GetCurrentFormat() sunTimeFormat {

	return currentSunTime.currentFormat
}

//
func (currentSunTime *SunTime) GetCurrentDateSeparator() string {

	return currentSunTime.currentDateSeparator
}

//
func (currentSunTime *SunTime) GetSunTimeAsTimestamp() int64 {

	return currentSunTime.asTimestamp
}

//
func (currentSunTime *SunTime) SetCurrentDateSeparator(newSeparator string) {

	if currentSunTime.currentDateSeparator != newSeparator {
	
		currentSunTime.currentDateSeparator = newSeparator
		fmt.Printf(Green() + "Date separator updated successfully." + Reset() + "\n")
		
	} else {
	
		fmt.Printf(Red() + "Date separator already as " + newSeparator + "." + Reset() + "\n")
	
	}
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsDMYHMS() {

	if currentSunTime.currentFormat != DMYHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = timePresentationFunction(wishedTime.Day()) + currentSunTime.currentDateSeparator + wishedTime.Month().String() + currentSunTime.currentDateSeparator + timePresentationFunction(wishedTime.Year()) + " " + timePresentationFunction(wishedTime.Hour()) + ":" + timePresentationFunction(wishedTime.Minute()) + ":" + timePresentationFunction(wishedTime.Second())
		currentSunTime.currentFormat = DMYHMSFormat
		
		fmt.Printf(Green() + "Suntime converted in DMYHMS successfully." + Reset() + "\n")

        } else {

                fmt.Printf(Red() + "Suntime already in DMYHMS." + Reset() + "\n")
        }
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsYMDHMS() {

	if currentSunTime.currentFormat != YMDHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = timePresentationFunction(wishedTime.Year()) + currentSunTime.currentDateSeparator + wishedTime.Month().String() + currentSunTime.currentDateSeparator + timePresentationFunction(wishedTime.Day()) + " " + timePresentationFunction(wishedTime.Hour()) + ":" + timePresentationFunction(wishedTime.Minute()) + ":" + timePresentationFunction(wishedTime.Second())
		currentSunTime.currentFormat = YMDHMSFormat
		
		fmt.Printf(Green() + "Suntime converted in YMDHMS successfully." + Reset() + "\n")

        } else {

                fmt.Printf(Red() + "Suntime already in YMDHMS." + Reset() + "\n")
        }
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsMDYHMS() {

	if currentSunTime.currentFormat != MDYHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = wishedTime.Month().String() + currentSunTime.currentDateSeparator + timePresentationFunction(wishedTime.Day()) + currentSunTime.currentDateSeparator + timePresentationFunction(wishedTime.Year()) + " " + timePresentationFunction(wishedTime.Hour()) + ":" + timePresentationFunction(wishedTime.Minute()) + ":" + timePresentationFunction(wishedTime.Second())
		currentSunTime.currentFormat = MDYHMSFormat
		
		fmt.Printf(Green() + "Suntime converted in MDYHMS successfully." + Reset() + "\n")

	} else {

		fmt.Printf(Red() + "Suntime already in MDYHMS." + Reset() + "\n")
	}
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsTimestamp() {

	if currentSunTime.currentFormat != Timestamp {

		currentSunTime.sunTimeInCurrentFormat = strconv.FormatInt(currentSunTime.asTimestamp, 10)
		currentSunTime.currentFormat = Timestamp
		
		fmt.Printf(Green() + "Suntime converted in timestamp successfully." + Reset() + "\n")

	} else {

		fmt.Printf(Red() + "Suntime already in Timestamp." + Reset() + "\n")
	}
}
