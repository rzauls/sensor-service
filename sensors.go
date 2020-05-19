package main

import (
	"database/sql"
	"strconv"
)

type Sensor struct {
	SensorId int 		`json:"sensor_id"`
	Name string 		`json:"name"`
	SerialCode string   `json:"serial_code"`
	Data []DataPoint	`json:"data"`
}

type DataPoint struct {
	Name string 		`json:"metric_name"`
	Value float64 		`json:"rvalue"`
	Unit  string 		`json:"unit_name"`
	Precision int32 	`json:"precision"`
	Time string 		`json:"rtime"`
}

// GetAllSensors - fetch newest reading from every sensor
func GetAllSensors(db *sql.DB) ([]Sensor, error) {
	// fetch sensor IDs
	sensorIds, err := db.Query("SELECT sens.sensor_id FROM main.sensors sens")
	if err != nil {
		return nil, err
	}
	var data []Sensor
	for sensorIds.Next() {
		var id int
		err := sensorIds.Scan(&id)
		if err != nil {
			return nil, err
		}
		// fetch relevant data points for each sensor
		rows, err := db.Query(
			"SELECT sens.sensor_id, " +
			" sens.serial_code, " +
			" sens.name, " +
			" mes.rvalue, " +
			" u.unit_name, " +
			" met.metric_name, " +
			" u.\"precision\", " +
			" mes.rtime " +
			" FROM main.sensors sens " +
			" LEFT JOIN main.measures mes ON mes.sensor_id = sens.sensor_id " +
			" LEFT JOIN main.metrics met ON met.metric_id = mes.metric_id " +
			" LEFT JOIN main.units u ON u.unit_id = met.unit_id " +
			" WHERE sens.sensor_id = " + strconv.Itoa(id) +
			" GROUP BY met.metric_id " +
			" ORDER BY rtime DESC")
		if err != nil {
			return nil, err

		}
		var (
			sensorID   int
			serialCode string
			name       string
			rvalue     sql.NullFloat64
			unitName   sql.NullString
			metricName sql.NullString
			precision  sql.NullInt32
			rtime      sql.NullString
			dataPoints []DataPoint
		)
		for rows.Next() {
			err := rows.Scan(&sensorID, &serialCode, &name, &rvalue, &unitName, &metricName, &precision, &rtime)

			if err != nil {
				return nil, err
			}

			if rvalue.Valid {
				dataPoints = append(dataPoints, DataPoint{
					Name:      metricName.String,
					Value:     rvalue.Float64,
					Unit:      unitName.String,
					Precision: precision.Int32,
					Time:      rtime.String,
				})
			}
		}
		data = append(data, Sensor{
			SensorId: sensorID,
			Name:     name,
			SerialCode: serialCode,
			Data: dataPoints,
		})
	}
	return data, nil
}