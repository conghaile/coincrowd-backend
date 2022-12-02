package db

import "database/sql"

type PostgresStore struct {
	db *sql.DB
}
