package locations

import (
	"Groupie-tracker/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetLocations() structures.Locations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	location, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrLocations structures.Locations
	err = json.Unmarshal(location, &arrLocations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrLocations)
	return arrLocations
}
