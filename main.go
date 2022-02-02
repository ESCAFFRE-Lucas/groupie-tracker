package main

import (
	"fmt"
	"groupie-tracker/artists"
	"groupie-tracker/structures"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	data := artists.GetArtists()
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("artist-index.gohtml"))
	data := artists.GetArtists()
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", HomePage)
	server.HandleFunc("/artist", ArtistPage)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	fmt.Println("server listening on http://localhost:8000/")

	_ = http.ListenAndServe(":8000", server)
}

func IdInURL(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	id := request.URL.Path[len("/artist/"):]
	if strings.Contains(id, "/") {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("Tu hors de ma vue"))
	}
	_, err := writer.Write([]byte(id))
	if err != nil {
		fmt.Println(err)
	}
	intId, _ := strconv.Atoi(id)
	SpecificArtist := IdInJson(intId)
	tmpl := template.Must(template.ParseFiles("artist-index.gohtml"))
	data := SpecificArtist
	err = tmpl.Execute(writer, data)
	if err != nil {
		fmt.Println(err)
	}
}

func IdInJson(id int) structures.Artists {
	artist := artists.GetArtists()
	for i := range artist {
		if artist[i].Id == id {
			return artist[i]
		} else {
			continue
		}
	}
	return artist[0]
}
