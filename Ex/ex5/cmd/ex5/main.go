package main

import (
	"time"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Player struct {
	ID string `json:"uuid"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Birthday  time.Time    `json:"birthday"`
	Genre     string `json:"genre"`
	Sport     string `json:"sport"`
	TeamName  string `json:"team_name"`
}

type Players []Player

func main() {
	birthday1 := time.Date(1998, time.June, 26, 0, 0, 0, 0, time.UTC)
	birthday2 := time.Date(1988, time.January, 1, 0, 0, 0, 0, time.UTC)
	birthday3 := time.Date(1973, time.March, 26, 0, 0, 0, 0, time.UTC)
	birthday4 := time.Date(1984, time.December, 30, 0, 0, 0, 0, time.UTC)
	birthday5 := time.Date(1991, time.October, 10, 0, 0, 0, 0, time.UTC)
	birthday6 := time.Date(1988, time.November, 12, 0, 0, 0, 0, time.UTC)
	birthday7 := time.Date(1984, time.February, 12, 0, 0, 0, 0, time.UTC)
	birthday8 := time.Date(1986, time.February, 8, 0, 0, 0, 0, time.UTC)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go to /players")
	})

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {

		playerinfo := Players{
			Player{"47f1923a-0227-11ed-b939-0242ac120002","Ra", "TamaG", birthday1, "M", "cricket", "Caballeros de Kolkata"},
			Player{"04ef0cad-c366-486b-8a98-8480b7a57442","Fernando", "Uribe", birthday2, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
			Player{"47f1946a-0227-11ed-b939-0242ac120002","Imer", "Machado", birthday3, "M", "Futbol", "Naciopan"},
			Player{"6ac7fa78-9b91-48c6-b8cf-45e917844d4b","Legone", "James", birthday4, "M", "Basketball", "Lakers"},
			Player{"f37c12fd-e710-4d9b-99fa-fd4c60931b57","Mariana", "Pajon", birthday5, "F", "Bmx", "Colombia"},
			Player{"47f196cc-0227-11ed-b939-0242ac120002","Russell", "Wesbrick", birthday6, "M", "Basketball", "Lakers"},
			Player{"0ae0bdf1-2a75-405d-9470-9f525c9b597e","Caterine", "Ibargüen", birthday7, "F", "Triple Salto", "Colombia"},
			Player{"47f196cc-0227-11ed-b939-0242ac120002","Carlos", "Cow", birthday8, "M", "Futbol", "Club Deportivo Popular Junior Fútbol Club S.A."},
		}

		json.NewEncoder(w).Encode(playerinfo)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("error", err)
	}
}
