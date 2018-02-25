package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var sensorDataRepository = repository.SensorDataRepository{}

func init() {
	sensorDataRepository.Connect()
}

func InsertDataSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data domain.SensorData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	url := extractUrlSensor(r)
	idSensor, err := sensorRepository.FindIdSensorByUrl(url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	data.ID = bson.NewObjectId()
	data.Sensor = idSensor
	if err := sensorDataRepository.InsertSensorData(url, data); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})
}

func extractUrlSensor(r *http.Request) string {
	uri := r.RequestURI
	uri = strings.Replace(uri, "/", "", 1)
	return uri
}

func CreatingRoutesSensors(route *mux.Router) {

	sensors, err := sensorRepository.FindAllSensor()

	if err != nil {
		log.Fatal(err)
	}

	for _, sensor := range sensors {
		route.HandleFunc("/"+sensor.Url, InsertDataSensor).Methods("POST")
	}
	log.Println("Created Routes")
}
