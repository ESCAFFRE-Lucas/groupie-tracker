package main

import (
	"fmt"
	"groupie-tracker/artists"
	"html/template"
	"net/http"
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

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/")

	_ = http.ListenAndServe(":8000", server)
}
