package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/translator/internal/service"
	"github.com/alserov/translator/internal/validator"
	"net/http"
)

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
		valid:   validator.NewValidator(),
	}
}

type AuthHandler struct {
	service service.Service

	valid validator.Validator
}

func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ah.signup(w, r)
	case http.MethodPatch:
		ah.login(w, r)
	}
}

func (ah *AuthHandler) signup(w http.ResponseWriter, r *http.Request) {
	var user service.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, err)
		return
	}

	if err := ah.valid.Validate(user); err != nil {
		handleError(w, err)
		return
	}

	uuid, err := ah.service.CreateUser(r.Context(), user)
	if err != nil {
		fmt.Println(err)
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UuidResponse{Uuid: uuid})
}

func (ah *AuthHandler) login(w http.ResponseWriter, r *http.Request) {

}
