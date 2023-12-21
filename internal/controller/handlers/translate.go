package handlers

import (
	"github.com/alserov/translator/internal/service"
	"net/http"
)

func NewTranslatorHandler(service service.Service) *TranslatorHandler {
	return &TranslatorHandler{
		service: service,
	}
}

func (fth *TranslatorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fth.translate(w, r)
	}
}

type TranslatorHandler struct {
	service service.Service
}

func (fth *TranslatorHandler) translate(w http.ResponseWriter, r *http.Request) {
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

	translatedDocx, err := fth.service.TranslateDocx(r.Context(), service.TranslateParams{
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
