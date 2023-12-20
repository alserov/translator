package handlers

import (
	"code.sajari.com/docconv"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type TranslatorHandler struct {
}

func NewTranslatorHandler() *TranslatorHandler {
	return &TranslatorHandler{}
}

func (fth *TranslatorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fth.translate(w, r)
	}
}

func (fth *TranslatorHandler) translate(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	d, _, err := docconv.ConvertDocx(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	res, err := gt.Translate(d, from, to)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	f, _ := os.OpenFile("translate.txt", os.O_CREATE, 0644)
	defer os.Remove("translate.txt")
	defer f.Close()

	res = removeLinks(res)

	_, err = f.Write([]byte(res))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = f.WriteString(fmt.Sprintf("\nTranslated from %s to %s", from, to))

	txt, err := ioutil.ReadFile("translate.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=translated.docx")
	w.Write(txt)
}

func removeLinks(text string) string {
	words := strings.Split(text, " ")
	for i, w := range words {
		if strings.Contains(w, "http") || strings.Contains(w, "https") {
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
