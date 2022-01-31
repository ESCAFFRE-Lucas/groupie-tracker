package Relation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type index struct {
	Relations []relations `json:"index"`
}

type relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetRelations() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	relation, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrRelations index
	err = json.Unmarshal(relation, &arrRelations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrRelations)
}
