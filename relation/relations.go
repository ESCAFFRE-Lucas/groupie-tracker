package relation

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

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

	var arrRelations structures.Index
	err = json.Unmarshal(relation, &arrRelations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrRelations)
}
