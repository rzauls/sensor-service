package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := sql.Open("sqlite3", "./db/aranet.db")
	if err != nil {

		return err
	}
	rows, err := db.Query("SELECT sensor_id, serial_code, `name` FROM sensors")
	if err != nil {
		return err

	}
	var (
		sensorId int
		serialCode int
		name string
	)

	for rows.Next() {
		err := rows.Scan(&sensorId, &serialCode, &name)

		if err != nil {
			return err
		}

		fmt.Println(strconv.Itoa(sensorId) + " (" + strconv.Itoa(serialCode) + "): " + name)
	}
	return nil
}

