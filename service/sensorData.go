package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type ISensorDataService interface {
	Create(w http.ResponseWriter, r *http.Request)
	/*FssssssindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)*/
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

func extractUrlSensor(r *http.Request) string {
	uri := r.RequestURI
	uri = strings.Replace(uri, "/", "", 1)
	return uri
}
