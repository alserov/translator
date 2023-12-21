package mysql

import (
	"database/sql"
	"github.com/alserov/translator/internal/db/mysql/migrations"

	_ "github.com/go-sql-driver/mysql"
)

func MustConnect(dsn string) *sql.DB {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	migrations.MustMigrate(conn)

	return conn
}
