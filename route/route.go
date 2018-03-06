package route

import (
	"home-collect/servicecontainer"
	"log"
	"sync"

	"github.com/gorilla/mux"
)

type IMuxRouter interface {
	InitRouter() *mux.Router
}

type router struct{}

func (router *router) InitRouter() *mux.Router {
	serviceContainer := servicecontainer.ServiceContainer()
	userService := serviceContainer.InjectUserService()
	sensorService := serviceContainer.InjectSensorService()
	sensorDataService := serviceContainer.InjectSensorDataService()

	r := mux.NewRouter()
	r.HandleFunc("/users", userService.Create).Methods("POST")
	r.HandleFunc("/users", userService.FindAll).Methods("GET")
	r.HandleFunc("/users", userService.Delete).Methods("DELETE")
	r.HandleFunc("/users:id", userService.Update).Methods("PUT")
	r.HandleFunc("/users:id", userService.FindOne).Methods("GET")

	r.HandleFunc("/sensors", sensorService.Create).Methods("POST")
	r.HandleFunc("/sensors", sensorService.Delete).Methods("DELETE")
	r.HandleFunc("/sensors", sensorService.FindAll).Methods("GET")
	r.HandleFunc("/sensors:id", sensorService.Update).Methods("PUT")
	r.HandleFunc("/sensors:id", sensorService.FindOne).Methods("GET")

	urlRoutes, err := sensorService.GenerateUrlsSensor()
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range urlRoutes {
		r.HandleFunc(url, sensorDataService.Create).Methods("POST")
		r.HandleFunc(url, sensorDataService.FindAll).Methods("GET")
		r.HandleFunc(url+":id", sensorDataService.FindOne).Methods("GET")
	}

	log.Println("Routes created")
	return r
}

var (
	r          *router
	routerOnce sync.Once
)

func MuxRouter() IMuxRouter {
	if r == nil {
		routerOnce.Do(func() {
			r = &router{}
		})
	}

	return r
}
