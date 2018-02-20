package main

import (
	"home-collect/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/users", service.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":3000", route); err != nil {
		log.Fatal(err)
	}
}
