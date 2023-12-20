package controller

import "net/http"

type Handlers struct {
	Auth       http.Handler
	Translator http.Handler
}

func NewRouter(h *Handlers) {
	http.Handle("/auth", h.Auth)
	http.Handle("/translate", h.Translator)
}
