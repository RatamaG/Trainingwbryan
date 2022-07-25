package main

import (
	"log"
	"net/http"
	"github.com/RatamaG/docs/cmd/ex5-1/server"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", server.HomeHandler)
	r.HandleFunc("/players", server.PlayersHandler)
	r.HandleFunc("/players/{id}", server.IDHandler)
	r.HandleFunc("/update", server.UpdateHandler)
	r.HandleFunc("/update/{id}", server.UpdateIDHandler).Methods("PUT")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("error", err)
	}

}

