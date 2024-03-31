package database

import (
	"construction-organization-system/pkg/config"
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDB(dbConfig *config.DBConfig) (*sql.DB, error) {

	db, err := sql.Open("postgres", dbConfig.ConnectionInfo())
	if err != nil {
		return nil, err
	}

	return db, nil
}
