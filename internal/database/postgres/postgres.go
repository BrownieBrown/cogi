package postgres

import (
	"database/sql"
	"github.com/BrownieBrown/cogi.git/internal/config"
)

func Client(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.Postgres.Driver, config.Postgres.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
