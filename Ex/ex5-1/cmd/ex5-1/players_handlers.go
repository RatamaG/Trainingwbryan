package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Datas)
}
func IDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	for _, player := range Playerlist {
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

	var updatePlayer Player

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &updatePlayer)

	for i, Player := range Playerlist {
		if Player.ID == vars {
			Player.ID = updatePlayer.ID
			Player.FirstName = updatePlayer.FirstName
			Player.LastName = updatePlayer.LastName
			Player.Birthday = updatePlayer.Birthday
			Player.Genre = updatePlayer.Genre
			Player.TeamName = updatePlayer.TeamName
			Playerlist = append(Playerlist[:i], Player)

			json.NewEncoder(w).Encode(Player)
		}
	}
}