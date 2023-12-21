package db

import "time"

type User struct {
	Uuid            string
	Username        string
	Password        string
	Email           string
	RequestsBalance uint32
	CreatedAt       *time.Time
}
