package main

import (
	"fmt"
	"groupie-tracker/artists"
	"groupie-tracker/relation"
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

func RelationPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/relation.gohtml"))
	fmt.Println(r.URL.Path)
	groupRelation := relation.GetRelation(w, r)
	err := tmpl.Execute(w, groupRelation)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/artists", HomePage)       //Every artist
	server.HandleFunc("/artists/", ArtistPage)    //Single artist
	server.HandleFunc("/relation/", RelationPage) // every concert's dates and locations

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/artists")

	_ = http.ListenAndServe(":8000", server)
}
