package owmStructures

type OWMStruct struct {
	Coord Coordinates `json:"coord"`
	Weather [1]Weather `json:"weather"`
	Base string `json:"base"`
	Main Main `json:"main"`
	Visibility int `json:"visibility"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Sys Sys `json:"sys"`
	Timezone int `json:"timezone"`
	Id int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
	Message string `json:"message"`
}