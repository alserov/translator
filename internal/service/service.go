package service

import (
	"code.sajari.com/docconv"
	"context"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"io/ioutil"
	"os"
	"strings"
)

type Service interface {
	TranslateDocx(ctx context.Context, params TranslateParams, userUuid string) (translated []byte, err error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) TranslateDocx(ctx context.Context, params TranslateParams, userUuid string) (translated []byte, err error) {
	txt, _, err := docconv.ConvertDocx(params.File)
	if err != nil {
		return nil, err
	}

	if params.RemoveLinks {
		txt = removeLinks(txt)
	}

	res, err := gt.Translate(txt, params.From, params.To)
	if err != nil {
		return nil, err
	}

	f, _ := os.OpenFile("translate.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer os.Remove("translate.txt")
	defer f.Close()

	_, err = f.Write([]byte(res))
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString(fmt.Sprintf("\nTranslated from %s to %s", params.From, params.To))
	if err != nil {
		return nil, err
	}

	translated, err = ioutil.ReadFile("translate.txt")
	if err != nil {
		return nil, err
	}

	return translated, nil
}

func removeLinks(text string) string {
	words := strings.Split(text, " ")

	for i, w := range words {
		if strings.Contains(w, "http") {
			words = append(words[:i], words[i+1:]...)
		}
	}

	for i, w := range words {
		if strings.Contains(w, "https") {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return strings.Join(words, " ")
}
