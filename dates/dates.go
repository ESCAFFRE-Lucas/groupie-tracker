package dates

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type dates struct {
	Index []struct {
		Date []string `json:"dates"`
	}
}

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

	var arrDates dates
	err = json.Unmarshal(date, &arrDates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", arrDates)
}
