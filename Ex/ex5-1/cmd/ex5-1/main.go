package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Birthday  int    `json:"birthday"`
	Genre     string `json:"genre"`
	Sport     string `json:"sport"`
	TeamName  string `json:"team_name"`
}

type infoplayer []Player

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go to /players")
	})

	r.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		playerinfo := infoplayer{
			Player{"Ra", "TamaG", 30, "M", "cricket", "Caballeros de Kolkata"},
			Player{"Fernando", "Uribe", 34, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
			Player{"Imer", "Machado", 49, "M", "Futbol", "Naciopan"},
			Player{"Legone", "James", 37, "M", "Basketball", "Lakers"},
			Player{"Mariana", "Pajon", 30, "F", "Bmx", "Colombia"},
			Player{"Russell", "Wesbrick", 33, "M", "Basketball", "Lakers"},
			Player{"Caterine", "Ibargüen", 38, "F", "Triple Salto", "Colombia"},
			Player{"Carlos", "Bacca", 35, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
		}
		json.NewEncoder(w).Encode(playerinfo)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("errror", err)
	}

}
