package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Player struct {
	ID        string    `json:"ID"`
	FirstName string    `json:"name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
	Genre     string    `json:"genre"`
	Sport     string    `json:"sport"`
	TeamName  string    `json:"team_name"`
}

type Players []Player

var Playerlist = Players{

	{
		ID:        "47f1923a-0227-11ed-b939-0242ac120002",
		FirstName: "Ra",
		LastName:  "TamaG",
		Birthday:  time.Date(1998, time.June, 26, 12, 35, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Cricket",
		TeamName:  "Caballeros de Kolkata",
	},
	{
		ID:        "04ef0cad-c366-486b-8a98-8480b7a57442",
		FirstName: "Fernando",
		LastName:  "Uribe",
		Birthday:  time.Date(1988, time.January, 1, 11, 1, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Futbol",
		TeamName:  "Club Deportivo Popular Junior Fútbol Club S.A.",
	},
	{
		ID:        "47f1946a-0227-11ed-b939-0242ac120002",
		FirstName: "Imer",
		LastName:  "Machado",
		Birthday:  time.Date(1973, time.March, 26, 17, 16, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Futbol",
		TeamName:  "Naciopan",
	},
	{
		ID:        "6ac7fa78-9b91-48c6-b8cf-45e917844d4b",
		FirstName: "Legone",
		LastName:  "James",
		Birthday:  time.Date(1984, time.December, 30, 12, 9, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Basketball",
		TeamName:  "Lakers",
	},
	{
		ID:        "f37c12fd-e710-4d9b-99fa-fd4c60931b57",
		FirstName: "Mariana",
		LastName:  "Pajon",
		Birthday:  time.Date(1991, time.October, 10, 9, 12, 0, 0, time.UTC),
		Genre:     "F",
		Sport:     "BMX",
		TeamName:  "Colombia",
	},
	{
		ID:        "47f196cc-0227-11ed-b939-0242ac120002",
		FirstName: "Russell",
		LastName:  "Westbrick",
		Birthday:  time.Date(1988, time.November, 12, 18, 7, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Basketball",
		TeamName:  "Lakers",
	},
	{
		ID:        "0ae0bdf1-2a75-405d-9470-9f525c9b597e",
		FirstName: "Catherine",
		LastName:  "Ibarguen",
		Birthday:  time.Date(1984, time.February, 12, 19, 12, 0, 0, time.UTC),
		Genre:     "F",
		Sport:     "Salto Triple",
		TeamName:  "Colombia",
	},
	{
		ID:        "47f196cc-0227-11ed-b939-0242ac120002",
		FirstName: "Carlitos",
		LastName:  "Cow",
		Birthday:  time.Date(1998, time.June, 26, 9, 15, 0, 0, time.UTC),
		Genre:     "M",
		Sport:     "Futbol",
		TeamName:  "Club Deportivo Popular Junior Fútbol Club S.A.",
	},
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go to /players")
	})

	r.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Playerlist)
	})
	r.HandleFunc("/players/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)

		for _, player := range Playerlist {
			if player.ID == vars["id"] {
				json.NewEncoder(w).Encode(player)
			}
		}

	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("error", err)
	}

}
