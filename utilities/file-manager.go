package utilities

import (
	"Groupie-tracker/structures"
	"encoding/json"
	"io/ioutil"
)

const GroupieDataFile = "artists.json"

//The function below permit to save the progress of the game in a json file
func SaveInFile(data []structures.Artist) { //take and return the json encoding of the structure
	jsonData, _ := json.MarshalIndent(data, "", "") //return the json encoding of data either structure.Stock structure
	err := ioutil.WriteFile(GroupieDataFile, jsonData, 0644)
	if err != nil {
		return
	}
}
