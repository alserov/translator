package clickhouse

import (
	"github.com/alserov/translator/internal/db"
	"github.com/uptrace/go-clickhouse/ch"
)

type repository struct {
}

func NewRepository(db *ch.DB) db.Repository {
	return &repository{}
}

func (r repository) CreateUser(user db.User) (uuid string, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) TopUpBalance(userUuid string, reqAmount uint32) error {
	//TODO implement me
	panic("implement me")
}
