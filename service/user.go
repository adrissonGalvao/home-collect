package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

var userRepository = repository.UserRepository{}

func init() {
	userRepository.Connect()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = bson.NewObjectId()
	if err := userRepository.InsertUser(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})

}
func FindAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := userRepository.FindAllUser()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)
}
