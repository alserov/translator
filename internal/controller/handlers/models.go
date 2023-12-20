package handlers

import (
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type User struct {
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Email           string     `json:"email"`
	RequestsBalance uint32     `json:"requestsBalance"`
	CreatedAt       *time.Time `json:"createdAt"`
	Token           string     `json:"token"`
}

type GoogleTranslateReq struct {
	Text       []string `json:"q"`
	TargetLang string   `json:"target"`
}

type GoogleTranslateRes struct {
}
