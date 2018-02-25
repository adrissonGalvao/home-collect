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
	route.HandleFunc("/users", service.FindAllUser).Methods("GET")
	route.HandleFunc("/users", service.DeleteUser).Methods("DELETE")
	route.HandleFunc("/users:id", service.UpdateUser).Methods("PUT")
	route.HandleFunc("/users:id", service.FindOneUser).Methods("GET")

	route.HandleFunc("/sensors", service.CreateSensor).Methods("POST")
	route.HandleFunc("/sensors", service.DeleteSensor).Methods("DELETE")
	route.HandleFunc("/sensors", service.FindAllSensor).Methods("GET")
	route.HandleFunc("/sensors:id", service.UpdateSensor).Methods("PUT")
	route.HandleFunc("/sensors:id", service.FindOneSensor).Methods("GET")

	service.CreatingRoutesSensors(route)
	if err := http.ListenAndServe(":3000", route); err != nil {
		log.Fatal(err)
	}
}
