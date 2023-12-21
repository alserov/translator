package migrations

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MustMigrate(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic("failed to init db instance: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/mysql/migrations/", "mysql", driver)
	if err != nil {
		panic("failed to init  Migrate instance: " + err.Error())
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic("failed to migrate: " + err.Error())
	}
}
