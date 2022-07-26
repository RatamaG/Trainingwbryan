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
	r.HandleFunc("/players", server.GetPlayers)
	r.HandleFunc("/players/{id}", server.GetPlayer)
	r.HandleFunc("/update/{id}", server.UpdatePlayer).Methods("PUT")
	r.HandleFunc("/delete/{id}", server.DeletePlayer).Methods("DELETE")
	r.HandleFunc("/newplayer", server.CreatePlayer).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("error", err)
	}

}
