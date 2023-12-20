package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/translator/internal/db"
	"github.com/alserov/translator/internal/validator"
	"net/http"
)

type AuthHandler struct {
	repo  db.Repository
	valid validator.Validator
}

func NewAuthHandler(repo db.Repository) *AuthHandler {
	return &AuthHandler{
		repo: repo,
	}
}

func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ah.signup(w, r, ah.repo)
	case http.MethodPatch:
		ah.login(w, r, ah.repo)
	}
}

func (ah *AuthHandler) signup(w http.ResponseWriter, r *http.Request, repo db.Repository) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: fmt.Sprintf("failed to decode request body: %v", err),
		})
	}

	if err := ah.valid.Validate(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: fmt.Sprintf("invalid data: %v", err),
		})
	}

	fmt.Println(user)
}

func (ah *AuthHandler) login(w http.ResponseWriter, r *http.Request, repo db.Repository) {

}
