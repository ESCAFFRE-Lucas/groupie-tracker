package main

import (
	"Groupie-tracker/artists"
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	data := artists.GetArtists()
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/artist.gohtml"))
	fmt.Println(r.URL.Path)
	artist := artists.GetArtist(w, r)
	err := tmpl.Execute(w, artist)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/artist", HomePage)    //Every artist
	server.HandleFunc("/artist/", ArtistPage) //Single artist

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/artist")

	_ = http.ListenAndServe(":8000", server)
}
