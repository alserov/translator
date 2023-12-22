package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/translator/internal/service"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	if e, ok := err.(*service.Error); ok {
		w.WriteHeader(e.Code)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(err)
}


func NewTranslatorHandler(service service.Service) *TranslatorHandler {
	return &TranslatorHandler{
		service: service,
	}
}

func (th *TranslatorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		th.translate(w, r)
	}
}

type TranslatorHandler struct {
	service service.Service
}

func (th *TranslatorHandler) translate(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	links := r.URL.Query().Get("links")
	uuid := r.URL.Query().Get("uuid")

	var removeLinks bool
	if links == "true" {
		removeLinks = true
	}

	translatedDocx, err := th.service.TranslateDocx(r.Context(), service.TranslateParams{
		RemoveLinks: removeLinks,
		From:        from,
		To:          to,
		File:        file,
	}, uuid)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=translated.docx")
	w.Write(translatedDocx)
}
