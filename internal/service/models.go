package service

import (
	"io"
)

type Error struct {
	Err  string `json:"error"`
	Code int
}

func (e *Error) Error() string {
	return e.Err
}

type ServiceError struct {
	error
	ErrorType uint32
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
