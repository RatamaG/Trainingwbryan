package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/RatamaG/docs/cmd/ex5-1/functions"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Sr Gonazlo Martinez A. Go to /players")
}
func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(functions.Datas)
}
func IDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	for _, player := range functions.Playerlist {
		if player.ID == vars["id"] {
			json.NewEncoder(w).Encode(player)
		}
	}
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Sr Gonazlo Martinez A. Go to /update/{id} ")
}
func UpdateIDHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)["id"]

	var updatePlayer functions.Player

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &updatePlayer)

	for i, Player := range functions.Playerlist {
		if Player.ID == vars {
			Player.ID = updatePlayer.ID
			Player.FirstName = updatePlayer.FirstName
			Player.LastName = updatePlayer.LastName
			Player.Birthday = updatePlayer.Birthday
			Player.Genre = updatePlayer.Genre
			Player.TeamName = updatePlayer.TeamName
			functions.Playerlist = append(functions.Playerlist[:i], Player)

			json.NewEncoder(w).Encode(Player)
		}
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/players", PlayersHandler)
	r.HandleFunc("/players/{id}", IDHandler)
	r.HandleFunc("/update", UpdateHandler)
	r.HandleFunc("/update/{id}", UpdateIDHandler).Methods("PUT")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("error", err)
	}

}
