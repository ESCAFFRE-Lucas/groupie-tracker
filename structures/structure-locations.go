package structures

type Locations struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`

	Dates string `json:"dates"`
}
