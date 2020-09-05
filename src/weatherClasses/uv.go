package weatherClasses

//
type UVRisk int

//
const (

	Low = iota
	Moderate
	High
	Very_High
	Extreme
)

//
func (risk UVRisk) String() string {

	return [...]string{"Low", "Moderate", "High", "Very high", "Extreme"}[risk]
}

//
type UV struct {

	//
	index int64
	risk UVRisk
}

//
func (currentUV *UV) InitializeUV(value int64) {

	var determiedUVRisk UVRisk

	if value <= 2 {

		determiedUVRisk = Low

	} else if 3 <= value && value <= 5 {

		determiedUVRisk = Moderate

	} else if 6 <= value && value <= 7 {

		determiedUVRisk = High

	} else if 8 <= value && value <= 10 {

		determiedUVRisk = Very_High

	} else {

		determiedUVRisk = Extreme

	}

	currentUV.index = value
	currentUV.risk = determiedUVRisk
}

//
func (currentUV *UV) GetIndex() int64 {

	return currentUV.index
}

//
func (currentUV *UV) GetRisk() UVRisk {

	return currentUV.risk
}
