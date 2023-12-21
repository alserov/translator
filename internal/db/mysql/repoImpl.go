package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/translator/internal/db"
)

func NewRepository(db *sql.DB) db.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r repository) DebitRequests(ctx context.Context, userUuid string, reqAmount uint32) error {
	query := `UPDATE users SET requests_balance = requests_balance - ? WHERE uuid = ?`

	_, err := r.db.Exec(query, reqAmount, userUuid)
	if err != nil {
		return errors.New("failed to execute query: " + err.Error())
	}

	return nil
}

func (r repository) CreateUser(ctx context.Context, user db.User) (err error) {
	query := `INSERT INTO users (username,password, email, requests_balance,created_at) VALUES (?,?,?,?,?)`

	_, err = r.db.Exec(query, user.Username, user.Password, user.Email, user.RequestsBalance, user.CreatedAt)
	if err != nil {
		return errors.New("failed to execute query: " + err.Error())
	}

	return nil
}

func (r repository) TopUpRequests(ctx context.Context, userUuid string, reqAmount uint32) error {
	//TODO implement me
	panic("implement me")
}
