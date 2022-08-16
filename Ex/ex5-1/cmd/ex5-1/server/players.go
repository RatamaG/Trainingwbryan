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
		if Player.ID.String() == vars["id"] {
			Response := pkg.Response{
				Status: http.StatusOK,
				Data:   pkg.Players{Player},
			}
			json.NewEncoder(w).Encode(Response)
		}
		fmt.Println("Este es mi player.ID.String()", Player.ID.String())
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

	isValid, message := updatePlayer.Validate()

	if !isValid {
		response := pkg.Response{
			Status:  http.StatusUnprocessableEntity,
			Data:    pkg.Players{},
			Message: message,
		}

		json.NewEncoder(w).Encode(response)
	}

	for i, Player := range pkg.Playerlist {
		if Player.ID.String() == vars {
			Player.ID = updatePlayer.ID
			Player.FirstName = updatePlayer.FirstName
			Player.LastName = updatePlayer.LastName
			Player.Birthday = updatePlayer.Birthday
			Player.Genre = updatePlayer.Genre
			Player.TeamName = updatePlayer.TeamName
			pkg.Playerlist = append(pkg.Playerlist[:i], Player)

			Response := pkg.Response{
				Status:  http.StatusOK,
				Data:    pkg.Playerlist,
				Message: "Player Updated successfully",
			}
			json.NewEncoder(w).Encode(Response)
		}

	}
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {

	var DeletePlayer pkg.Player

	vars := mux.Vars(r)["id"]
	
	isValid, message := DeletePlayer.Validate()

	if !isValid {
		response := pkg.Response{
			Status:  http.StatusUnprocessableEntity,
			Data:    pkg.Players{},
			Message: message,
		}

		json.NewEncoder(w).Encode(response)
	}

	for i, Player := range pkg.Playerlist {
		if Player.ID.String() == vars {
			pkg.Playerlist = append(pkg.Playerlist[:i], pkg.Playerlist[i+1:]...)

			Response := pkg.Response{
				Status:  http.StatusOK,
				Data:    pkg.Playerlist,
				Message: "Player Delete successfully",
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

	isValid, message := newPlayer.Validate()

	if !isValid {
		response := pkg.Response{
			Status:  http.StatusUnprocessableEntity,
			Data:    pkg.Players{},
			Message: message,
		}

		json.NewEncoder(w).Encode(response)
	}
	pkg.Playerlist = append(pkg.Playerlist, newPlayer)

	Response := pkg.Response{
		Status:  http.StatusOK,
		Data:    pkg.Playerlist,
		Message: "Player Created successfully",
	}
	json.NewEncoder(w).Encode(Response)

}
