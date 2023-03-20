package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func (p *Pool) Migrate() error {
	driver, err := sqlite3.WithInstance(p.db.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite3", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no migration changes made")
		} else {
			return err
		}
	}

	return nil
}
