package locations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type locations struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`

	Dates string `json:"dates"`
}

func GetLocations() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	location, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrLocations locations
	err = json.Unmarshal(location, &arrLocations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrLocations)
}
