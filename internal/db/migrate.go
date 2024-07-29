package db

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/database/postgres"
)

func Migrate(url string) (err error) {
	migrations, err := migrate.New("file://migrations/postgres", url)
	if err != nil {
		return
	}
	if err = migrations.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return
	}
	return
}
