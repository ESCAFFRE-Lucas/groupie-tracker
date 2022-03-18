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

//This func permit to get informations for all artists in the API
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

//This func permit to get informations for one specific artist in the API by using his id
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

//Same as GetArtist, but for lofi (can't make 2 same requests)
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

//Get a specific artist's id, for the GetArtistForLofi func
func GetArtistsId(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/artists/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Error"))
		return ""
	}
	return id
}

//Get a specific artist's id, for the GetArtist func
func GetArtistsIdForLofi(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/lofi/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Error"))
		return ""
	}
	return id
}
