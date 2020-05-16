package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./db/aranet.db")
	rows, err := database.Query("SELECT sensor_id, serial_code, `name` FROM sensors")
	if err != nil {
		panic(err.Error())
	}
	var sensorId int
	var serialCode int
	var name string
	for rows.Next() {
		err := rows.Scan(&sensorId, &serialCode, &name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(strconv.Itoa(sensorId) + " (" + strconv.Itoa(serialCode) + "): " + name)
	}
}