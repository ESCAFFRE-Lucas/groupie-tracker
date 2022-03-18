package main

import (
	"fmt"
	"groupie-tracker/artists"
	"groupie-tracker/relation"
	"html/template"
	"net/http"
)

//Setup the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := artists.GetArtists()
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

//Setup the artit page
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/artistInfo.gohtml"))
	fmt.Println(r.URL.Path)
	artist := artists.GetArtist(w, r)
	err := tmpl.Execute(w, artist)
	if err != nil {
		fmt.Println(err)
	}
}

//Setup the relation page
func RelationPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/relation.gohtml"))
	fmt.Println(r.URL.Path)
	groupRelation := relation.GetRelation(w, r)
	err := tmpl.Execute(w, groupRelation)
	if err != nil {
		fmt.Println(err)
	}
}

//Setup the about us page
func AboutusPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/aboutus.gohtml"))
	fmt.Println(r.URL.Path)
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

//Setup the lofi page
func LofiPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/lofi.gohtml"))
	fmt.Println(r.URL.Path)
	artist := artists.GetArtistForLofi(w, r)
	err := tmpl.Execute(w, artist)
	if err != nil {
		fmt.Println(err)
	}
}

//Start the server and handle all the pages
func main() {
	server := http.NewServeMux()

	server.HandleFunc("/artists", HomePage)       //Every artist
	server.HandleFunc("/artists/", ArtistPage)    //Single artist
	server.HandleFunc("/relation/", RelationPage) // every concert's dates and locations
	server.HandleFunc("/aboutus", AboutusPage)    //About us page
	server.HandleFunc("/lofi/", LofiPage)         //Relaxation page

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/artists")

	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println(err)
	}
}
