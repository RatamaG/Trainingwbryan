package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

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
