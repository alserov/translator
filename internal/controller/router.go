package controller

import "net/http"

type Handlers struct {
	Translator http.Handler
}

func NewRouter(h *Handlers) {
	http.Handle("/translate", h.Translator)
}
