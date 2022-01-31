package Relation

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type relations struct {
	Index []struct {
		DatesLocations []struct {
			Locations []string
		} `json:"datesLocations"`
	} `json:"index"`
}

func GetRelations() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	relation, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrRelations relations
	err = json.Unmarshal(relation, &arrRelations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrRelations)
}
