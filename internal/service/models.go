package service

import (
	"io"
	"time"
)

type Error struct {
	Err  string `json:"error"`
	Code int
}

func (e *Error) Error() string {
	return e.Err
}

const (
	ServerError = iota
	UserError
)

type ServiceError struct {
	error
	ErrorType uint32
}

type User struct {
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Email           string     `json:"email"`
	RequestsBalance uint32     `json:"requestsBalance"`
	CreatedAt       *time.Time `json:"createdAt"`
	Token           string     `json:"token"`
}

type TranslateParams struct {
	File io.Reader

	From        string
	To          string
	RemoveLinks bool
}

type GoogleTranslateReq struct {
	Text       []string `json:"q"`
	TargetLang string   `json:"target"`
}

type GoogleTranslateRes struct {
}
