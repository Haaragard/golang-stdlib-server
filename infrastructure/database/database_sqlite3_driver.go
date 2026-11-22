package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseSqlite3Driver struct {
	dsn  string
	pool *sql.DB
}

func NewDatabaseSqlite3Driver(dsn string) *DatabaseSqlite3Driver {
	return &DatabaseSqlite3Driver{dsn: dsn}
}

func (db *DatabaseSqlite3Driver) Open() error {
	pool, err := sql.Open("sqlite3", db.dsn)
	if err != nil {
		return err
	}
	db.pool = pool
	return nil
}

func (db *DatabaseSqlite3Driver) Close() {
	db.pool.Close()
}

func (db *DatabaseSqlite3Driver) IsOpen() bool {
	return db.pool != nil
}

func (db *DatabaseSqlite3Driver) Get() {
	fmt.Printf("Query made\n")
}
