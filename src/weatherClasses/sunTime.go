package weatherClasses

//
import (
	"fmt"
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
type SunTime struct {

	//
	sunTimeInCurrentFormat string
	currentFormat sunTimeFormat
	asTimestamp int64
}

//
func (currentSunTime *SunTime) InitializeSunTime(value int64) {

	currentSunTime.sunTimeInCurrentFormat = strconv.FormatInt(value, 10)
	currentSunTime.currentFormat = Timestamp
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
func (currentSunTime *SunTime) GetSunTimeAsTimestamp() int64 {

	return currentSunTime.asTimestamp
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsTimestamp() {

	if currentSunTime.currentFormat != Timestamp {

		currentSunTime.sunTimeInCurrentFormat = strconv.FormatInt(currentSunTime.asTimestamp, 10)
		currentSunTime.currentFormat = Timestamp

	} else {

		fmt.Printf(Red() + "Suntime already in Timestamp." + Reset() + "\n")
	}
}
