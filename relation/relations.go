package relation

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structures"
	"io/ioutil"
	"net/http"
	"strings"
)

//This func permit to get informations of relation in the API by using his id
func GetRelation(w http.ResponseWriter, r *http.Request) structures.Relations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + GetRelationsId(w, r))
	if err != nil {
		fmt.Println(err)
	}

	strRelations, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var relation structures.Relations
	fmt.Println(relation)
	err = json.Unmarshal(strRelations, &relation)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(relation)
	return relation
}

//Get a specific relation's id, for the GetRelation func
func GetRelationsId(w http.ResponseWriter, r *http.Request) string {
	id := r.URL.Path[len("/relation/"):]
	if strings.Contains(id, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("don't go here man..."))
		return ""
	}
	return id
}
