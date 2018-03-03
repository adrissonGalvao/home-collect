package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

type ISensorDataService interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
}
type SensorDataService struct {
	repository.ISensorDataRepository
	repository.ISensorRepository
}

func (sd *SensorDataService) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data domain.SensorData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	url := extractUrlSensor(r)
	idSensor, err := sd.FindIdSensorByUrl(url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	data.ID = bson.NewObjectId()
	data.Sensor = idSensor
	if err := sd.InsertSensorData(url, data); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})
}

func (sd *SensorDataService) FindAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	url := extractUrlSensor(r)
	sensorsData, err := sd.FindAllSensorData(url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, sensorsData)
}

func (sd *SensorDataService) FindOne(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	url := extractUrlSensor(r)
	sensorData, err := sd.FindByIdSendorData(params["id"], url)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensorData)
}
func extractUrlSensor(r *http.Request) string {
	uri := r.RequestURI
	uri = strings.Replace(uri, "/", "", 1)
	return uri
}
