package structures

type Index struct {
	Relations []relations `json:"index"`
}

type relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
