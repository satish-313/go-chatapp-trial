package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	dp, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/gochatapp?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &Database{db: dp}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
