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

func GetArtist(w http.ResponseWriter, r *http.Request) structures.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + GetArtistsId(w, r))
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

func GetArtistForLofi(w http.ResponseWriter, r *http.Request) structures.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + GetArtistsIdForLofi(w, r))
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

func GetArtistsId(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/artists/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Error"))
		return ""
	}
	return id
}

func GetArtistsIdForLofi(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/lofi/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Error"))
		return ""
	}
	return id
}
