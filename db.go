package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB - exported DB handler
var DB *sql.DB

// InitDB - initialize and open DB connection
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./db/aranet.db")
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err

	}
	return nil
}
