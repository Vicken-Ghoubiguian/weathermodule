package owmStructures

type OWMStruct struct {
	Coord Coordinates `json:"coord"`
	Weather [1]Weather `json:"weather"`
	Base string `json:"base"`
	Main Main `json:"main"`
	Visibility int64 `json:"visibility"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Dt int64 `json:"dt"`
	Sys Sys `json:"sys"`
	Timezone int64 `json:"timezone"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Cod int64 `json:"cod"`
	Message string `json:"message"`
}
