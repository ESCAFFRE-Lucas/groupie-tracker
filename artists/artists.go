package artists

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"groupie-tracker/utilities"
	"io/ioutil"
	"log"
	"net/http"
)

func GetArtists() []structures.Artists {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	names, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrArtists []structures.Artists
	err = json.Unmarshal(names, &arrArtists)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arrArtists)
	utilities.SaveInFile(arrArtists)
	return arrArtists
}
