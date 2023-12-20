package clickhouse

import (
	"github.com/uptrace/go-clickhouse/ch"
)

func MustConnect(dsn string) *ch.DB {
	conn := ch.Connect(ch.WithDSN(dsn))
	return conn
}
