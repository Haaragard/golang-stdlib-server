package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseSqlite3Driver struct {
	pool *sql.DB
}

func NewDatabaseSqlite3Driver(dsn string) (*DatabaseSqlite3Driver, error) {
	pool, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("Could not open sqlite3 database with error: %w", err)
	}
	return &DatabaseSqlite3Driver{pool: pool}, nil
}

func (db *DatabaseSqlite3Driver) Get() {
	fmt.Printf("Query made\n")
}

func (db *DatabaseSqlite3Driver) Close() {
	db.pool.Close()
}
