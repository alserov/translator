package db

import "time"

type User struct {
	Username        string
	Password        string
	Email           string
	RequestsBalance uint32
	CreatedAt       *time.Time
}
