package main

import (
	"fmt"
	"groupie-tracker/Artists"
	"groupie-tracker/Relation"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	Relation.GetRelations()
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	data := Artists.GetArtists()
	_ = tmpl.Execute(w, data)

}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", HomePage)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/")

	_ = http.ListenAndServe(":8000", server)
}
