package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB - initialize and open DB connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/aranet.db")
	if err != nil {

		return nil, err
	}
	rows, err := db.Query("SELECT sensor_id, serial_code, `name` FROM sensors")
	if err != nil {
		return nil, err

	}
	var (
		sensorID   int
		serialCode int
		name       string
	)

	for rows.Next() {
		err := rows.Scan(&sensorID, &serialCode, &name)

		if err != nil {
			return nil, err
		}

		fmt.Println(strconv.Itoa(sensorID) + " (" + strconv.Itoa(serialCode) + "): " + name)
	}

	return db, nil
}
