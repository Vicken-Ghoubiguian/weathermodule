package weatherClasses

//
type Humidity struct {

	value int64
	unitScale string
}

//
func (currentHumidity *Humidity) InitializeHumidity(humidityValue int64) {

	currentHumidity.value = humidityValue
	currentHumidity.unitScale = "%%"
}

//
func (currentHumidity *Humidity) GetHumidityValue() int64 {

	return currentHumidity.value
}

//
func (currentHumidity *Humidity) GetHumidityUnitScale() string {

	return currentHumidity.unitScale
}
