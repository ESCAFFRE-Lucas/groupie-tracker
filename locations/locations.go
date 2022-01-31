package locations

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io/ioutil"
	"log"
	"net/http"
)

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

	var arrLocations structures.Locations
	err = json.Unmarshal(location, &arrLocations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrLocations)
}
