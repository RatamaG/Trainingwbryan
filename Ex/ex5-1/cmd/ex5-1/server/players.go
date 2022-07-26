package server

import (
	"encoding/json"
	"fmt"
	"github.com/RatamaG/docs/cmd/ex5-1/pkg"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	Response := pkg.Response{
		Status: http.StatusOK,
		Data:   pkg.Playerlist,
	}
	json.NewEncoder(w).Encode(Response)
}
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, Player := range pkg.Playerlist {
		if Player.ID == vars["id"] {
			json.NewEncoder(w).Encode(Player)
			Response := pkg.Response{
				Status: http.StatusOK,
				Data:   pkg.Players{Player},
			}
			json.NewEncoder(w).Encode(Response)
		}

	}
}
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["id"]

	var updatePlayer pkg.Player

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &updatePlayer)

	for i, Player := range pkg.Playerlist {
		if Player.ID == vars {
			Player.ID = updatePlayer.ID
			Player.FirstName = updatePlayer.FirstName
			Player.LastName = updatePlayer.LastName
			Player.Birthday = updatePlayer.Birthday
			Player.Genre = updatePlayer.Genre
			Player.TeamName = updatePlayer.TeamName
			pkg.Playerlist = append(pkg.Playerlist[:i], Player)

			json.NewEncoder(w).Encode(Player)
		}
	}
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["id"]

	for i, Player := range pkg.Playerlist {
		if Player.ID == vars {
			pkg.Playerlist = append(pkg.Playerlist[:i], pkg.Playerlist[i+1:]...)

			Response := pkg.Response{
				Status: http.StatusOK,
				Data:   pkg.Playerlist,
			}
			json.NewEncoder(w).Encode(Response)
		}
	}
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {

	var newPlayer pkg.Player

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &newPlayer)
	pkg.Playerlist = append(pkg.Playerlist, newPlayer)

	Response := pkg.Response{
		Status: http.StatusOK,
		Data:   pkg.Playerlist,
	}
	json.NewEncoder(w).Encode(Response)
}
