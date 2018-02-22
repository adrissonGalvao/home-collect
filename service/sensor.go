package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

var sensorRepository = repository.SensorRepository{}

func init() {
	sensorRepository.Connect()
}

func CreateSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sensor.ID = bson.NewObjectId()
	if err := sensorRepository.InsertSensor(sensor); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})
}

func FindAllSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sensors, err := sensorRepository.FindAllSensor()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensors)
}

func FindOneSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	sensor, err := sensorRepository.FindByIdSensor(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensor)
}

func UpdateSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error())
		return
	}
	if err := sensorRepository.UpdateSensor(sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, sensor)
}

func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := sensorRepository.DeleteSensor(sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensor)
}
