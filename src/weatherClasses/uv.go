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
	index int
	risk UVRisk
}

//
func InitializeUV(value int) *UV {

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

	return &UV{index: value, risk: determiedUVRisk}
}

//
func (currentUV *UV) GetIndex() int {

	return currentUV.index
}

//
func (currentUV *UV) GetRisk() UVRisk {

	return currentUV.risk
}
