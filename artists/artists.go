package artists

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"groupie-tracker/utilities"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetArtists() []structures.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	names, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var arrArtists []structures.Artist
	err = json.Unmarshal(names, &arrArtists)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arrArtists)
	utilities.SaveInFile(arrArtists)
	return arrArtists
}

func GetArtist() structures.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/1")
	if err != nil {
		fmt.Println(err)
	}

	strArtist, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var artist structures.Artist
	err = json.Unmarshal(strArtist, &artist)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(artist)
	return artist
}

func GetArtistsById(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/artist/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("degage en*****"))
		return ""
	}
	return id
}
