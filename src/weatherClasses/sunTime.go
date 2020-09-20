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
func (currentSunTime *SunTime) SetCurrentFormatAsDMYHMS() {

	if currentSunTime.currentFormat != DMYHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = strconv.Itoa(wishedTime.Day()) + "/" + wishedTime.Month().String() + "/" + strconv.Itoa(wishedTime.Year()) + " " + strconv.Itoa(wishedTime.Hour()) + ":" + strconv.Itoa(wishedTime.Minute()) + ":" + strconv.Itoa(wishedTime.Second())
		currentSunTime.currentFormat = DMYHMSFormat

        } else {

                fmt.Printf(Red() + "Suntime already in DMYHMS." + Reset() + "\n")
        }
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsYMDHMS() {

	if currentSunTime.currentFormat != YMDHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = strconv.Itoa(wishedTime.Year()) + "/" + wishedTime.Month().String() + "/" + strconv.Itoa(wishedTime.Day()) + " " + strconv.Itoa(wishedTime.Hour()) + ":" + strconv.Itoa(wishedTime.Minute()) + ":" + strconv.Itoa(wishedTime.Second())
		currentSunTime.currentFormat = YMDHMSFormat

        } else {

                fmt.Printf(Red() + "Suntime already in YMDHMS." + Reset() + "\n")
        }
}

//
func (currentSunTime *SunTime) SetCurrentFormatAsMDYHMS() {

	if currentSunTime.currentFormat != MDYHMSFormat {

		wishedTime := time.Unix(currentSunTime.asTimestamp, 0)
		currentSunTime.sunTimeInCurrentFormat = wishedTime.Month().String() + "/" + strconv.Itoa(wishedTime.Day()) + "/" + strconv.Itoa(wishedTime.Year()) + " " + strconv.Itoa(wishedTime.Hour()) + ":" + strconv.Itoa(wishedTime.Minute()) + ":" + strconv.Itoa(wishedTime.Second())
		currentSunTime.currentFormat = MDYHMSFormat

	} else {

		fmt.Printf(Red() + "Suntime already in MDYHMS." + Reset() + "\n")
	}
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
