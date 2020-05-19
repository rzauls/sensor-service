package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Sensor struct {
	Name string `json:"name"`
	Val1 int    `json:"val1"`
	Val2 int    `json:"val2"`
}

type Sensors []Sensor


// GetAllSensors - fetch newest reading from every sensor
func GetAllSensors(db *sql.DB) (Sensors, error) {
	rows, err := db.Query(
		"SELECT * FROM main.sensors sens " +
			"LEFT JOIN main.measures mes ON mes.sensor_id = sens.sensor_id " +
			"LEFT JOIN main.metrics met ON met.metric_id = mes.metric_id " +
			"LEFT JOIN main.units u ON u.unit_id = met.unit_id " +
			"WHERE sens.sensor_id = 22")
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
	sensors := Sensors{
		Sensor{
			Name: "test",
			Val1: 0,
			Val2: 1,
		},
	}
	return sensors, nil
}