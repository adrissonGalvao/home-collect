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

type ISensorService interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	VerifyIntegrityUrlSensor(id string) bool
	GenerateUrlsSensor() ([]string, error)
}
type SensorService struct {
	repository.ISensorRepository
}

func (s *SensorService) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sensor.Url = strings.ToLower(sensor.Url)
	if statusUrl := s.VerifyIntegrityUrlSensor(sensor.Url); statusUrl != false {
		sensor.ID = bson.NewObjectId()
		if err := s.InsertSensor(sensor); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request")
			return
		}
		respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})
	} else {
		respondWithError(w, http.StatusBadRequest, "Not available url")
	}

}

func (s *SensorService) FindAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sensors, err := s.FindAllSensor()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensors)
}

func (s *SensorService) FindOne(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	sensor, err := s.FindByIdSensor(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensor)
}

func (s *SensorService) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error())
		return
	}
	if err := s.UpdateSensor(sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensor)
}

func (s *SensorService) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var sensor domain.Sensor

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.DeleteSensor(sensor); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, sensor)
}

func (s *SensorService) VerifyIntegrityUrlSensor(id string) bool {
	statusUrl, err := s.FindUrlSensor(id)
	if err != nil {
		return false
	}
	return statusUrl
}
func (s *SensorService) GenerateUrlsSensor() ([]string, error) {

	sensors, err := s.FindAllSensor()
	var sensorUrls []string
	for _, sensor := range sensors {
		sensorUrls = append(sensorUrls, "/"+sensor.Url)

	}
	return sensorUrls, err
}
