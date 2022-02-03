package main

import (
	"fmt"
	"groupie-tracker/artists"
	"html/template"
	"net/http"
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
	fmt.Println(r.URL.Path)
	id := r.URL.Path[len("/artist/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("degage en*****"))
		return
	}
	_, err := w.Write([]byte(id))
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, id)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/artist", HomePage)
	server.HandleFunc("/artist/", ArtistPage)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/artist")

	_ = http.ListenAndServe(":8000", server)
}
