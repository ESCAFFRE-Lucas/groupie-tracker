package dates

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetDates() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	date, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrDates structures.Dates
	err = json.Unmarshal(date, &arrDates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrDates)
}
