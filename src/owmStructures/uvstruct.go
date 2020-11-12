package owmStructures

type UVStruct struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Date_iso string `json:"date_iso"`
	Date int64 `json:"date"`
	Value float64 `json:"value"`
}
