package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Player struct {
	Id string `json:"uuid"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Birthday  int    `json:"birthday"`
	Genre     string `json:"genre"`
	Sport     string `json:"sport"`
	TeamName  string `json:"team_name"`
}

type infoplayer []Player

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go to /players")
	})

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {

		playerinfo := infoplayer{
			Player{"47f1923a-0227-11ed-b939-0242ac120002","Ra", "TamaG", 30, "M", "cricket", "Caballeros de Kolkata"},
			Player{"04ef0cad-c366-486b-8a98-8480b7a57442","Fernando", "Uribe", 34, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
			Player{"47f1946a-0227-11ed-b939-0242ac120002","Imer", "Machado", 49, "M", "Futbol", "Naciopan"},
			Player{"6ac7fa78-9b91-48c6-b8cf-45e917844d4b","Legone", "James", 37, "M", "Basketball", "Lakers"},
			Player{"f37c12fd-e710-4d9b-99fa-fd4c60931b57","Mariana", "Pajon", 30, "F", "Bmx", "Colombia"},
			Player{"47f196cc-0227-11ed-b939-0242ac120002","Russell", "Wesbrick", 33, "M", "Basketball", "Lakers"},
			Player{"0ae0bdf1-2a75-405d-9470-9f525c9b597e","Caterine", "Ibargüen", 38, "F", "Triple Salto", "Colombia"},
			Player{"47f196cc-0227-11ed-b939-0242ac120002","Carlos", "Bacca", 35, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
		}

		json.NewEncoder(w).Encode(playerinfo)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("error", err)
	}
}
