package main

import (
	"groupie-tracker/Artists"
	"groupie-tracker/Dates"
	"groupie-tracker/Relation"
)

func main() {
	Artists.GetArtists()
	Dates.GetDates()
	Relation.GetRelations()
}
