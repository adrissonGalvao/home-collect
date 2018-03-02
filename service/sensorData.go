package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type ISersoDataService interface {
	Create(w http.ResponseWriter, r *http.Request)
	/*FindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)*/
}
type SersoDataService struct {
	repository.ISensorDataRepository
	repository.ISensorRepository
}

func (sd *SersoDataService) Create(w http.ResponseWriter, r *http.Request) {
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
