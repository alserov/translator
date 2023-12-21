package service

import (
	"code.sajari.com/docconv"
	"context"
	"fmt"
	"github.com/alserov/translator/internal/db"
	gt "github.com/bas24/googletranslatefree"
	id "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Service interface {
	CreateUser(ctx context.Context, user User) (uuid string, err error)
	TranslateDocx(ctx context.Context, params TranslateParams, userUuid string) (translated []byte, err error)
}

func NewService(repo db.Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo db.Repository
}

const (
	DEFAULT_REQUEST_BALANCE = 50
)

func (s *service) CreateUser(ctx context.Context, user User) (uuid string, err error) {
	uuid = id.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	createdAt := time.Now()

	err = s.repo.CreateUser(ctx, db.User{
		Uuid:            uuid,
		Username:        user.Username,
		Password:        string(hashedPassword),
		Email:           user.Email,
		RequestsBalance: DEFAULT_REQUEST_BALANCE,
		CreatedAt:       &createdAt,
	})
	if err != nil {
		return "", err
	}

	return uuid, nil
}

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

	f, _ := os.OpenFile("translate.txt", os.O_CREATE, 0644)
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

	if err = s.repo.DebitRequests(ctx, userUuid, 1); err != nil {
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
