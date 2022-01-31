package Artists

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type artists struct {
	Name string `json:"name"`
}

func GetArtists() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	names, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrArtists []artists
	err = json.Unmarshal(names, &arrArtists)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrArtists)
}
