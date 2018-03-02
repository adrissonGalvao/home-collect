package service

import (
	"encoding/json"
	"home-collect/domain"
	"home-collect/repository"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var userRepository = repository.UserRepository{}

type IUserService interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	repository.IUserRepository
}

func (us *UserService) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = bson.NewObjectId()
	if err := us.InsertUser(user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "sucess"})

}

func (us *UserService) FindAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	users, err := us.FindAllUser()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, users)
}

func (us *UserService) FindOne(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	user, err := us.FindByIdUser(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

func (us *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := us.DeleteUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, user)

}

func (us *UserService) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error())
		return
	}

	if err := us.UpdateUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, user)
}
