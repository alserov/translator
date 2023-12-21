package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/translator/internal/service"
	"net/http"
)

type UuidResponse struct {
	Uuid string `json:"uuid"`
}

func handleError(w http.ResponseWriter, err error) {
	if e, ok := err.(*service.Error); ok {
		w.WriteHeader(e.Code)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(err)
}
