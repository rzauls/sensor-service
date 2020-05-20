package main

import (
	"database/sql"
	"strconv"
)

type Sensor struct {
	SensorId int 			`json:"sensor_id"`
	Name string 			`json:"name"`
	SerialCode string   	`json:"serial_code"`
	Data []DataPoint		`json:"data"`
}

type DataPoint struct {
	Name string 			`json:"metric_name"`
	Value float64 			`json:"rvalue"`
	Unit  string 			`json:"unit_name"`
	Precision int32 		`json:"precision"`
	Time string 			`json:"rtime"`
}

type SensorMinMax struct {
	SensorId int 			`json:"sensor_id"`
	Name string 			`json:"name"`
	SerialCode string   	`json:"serial_code"`
	Data []DataPointMinMax 	`json:"data"`
}

type DataPointMinMax struct {
	Name string 			`json:"metric_name"`
	Unit  string 			`json:"unit_name"`
	Precision int32 		`json:"precision"`
	Min float64 			`json:"rvalue_min"`
	Max float64 			`json:"rvalue_max"`
	Time string 			`json:"date"`
}

// GetAnyNewestReading - fetch newest reading from every sensor
func GetAnyNewestReading(db *sql.DB) ([]Sensor, error) {
	// fetch sensor IDs
	sensorIds, err := db.Query("SELECT sens.sensor_id, sens.serial_code, sens.name FROM main.sensors sens")
	if err != nil {
		return nil, err
	}
	var data []Sensor
	for sensorIds.Next() {
		var (
			id int
			serialCode string
			name string
		)
		err := sensorIds.Scan(&id, &serialCode, &name)
		if err != nil {
			return nil, err
		}
		// fetch relevant data points for each sensor
		rows, err := db.Query(
			"SELECT mes.rvalue, " +
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
		// Null<T> to check for missing data
		var (
			rvalue     	sql.NullFloat64
			unitName   	sql.NullString
			metricName 	sql.NullString
			precision  	sql.NullInt32
			rtime		sql.NullString
			dataPoints 	[]DataPoint
		)
		for rows.Next() {
			err := rows.Scan(&rvalue, &unitName, &metricName, &precision, &rtime)

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
			SensorId: id,
			Name:     name,
			SerialCode: serialCode,
			Data: dataPoints,
		})
	}
	return data, nil
}

// GetAllSensorMinMaxOnDate - fetch min/max readings from every sensor on specific date
func GetAllSensorMinMaxOnDate(date string, db *sql.DB) ([]SensorMinMax, error) {
	// fetch sensor IDs
	sensorIds, err := db.Query("SELECT sens.sensor_id, sens.serial_code, sens.name FROM main.sensors sens")
	if err != nil {
		return nil, err
	}
	var data []SensorMinMax
	for sensorIds.Next() {
		var (
			id int
			serialCode string
			name string
		)
		err := sensorIds.Scan(&id, &serialCode, &name)
		if err != nil {
			return nil, err
		}
		// fetch relevant data points for each sensor
		rows, err := db.Query(
			"SELECT " +
			" MAX(mes.rvalue) AS rvalue_max, " +
			" MIN(mes.rvalue) as rvalue_min, " +
			" u.unit_name, " +
			" met.metric_name, " +
			" u.\"precision\", " +
			" mes.rtime " +
			" FROM main.sensors sens " +
			" LEFT JOIN main.measures mes ON mes.sensor_id = sens.sensor_id " +
			" LEFT JOIN main.metrics met ON met.metric_id = mes.metric_id " +
			" LEFT JOIN main.units u ON u.unit_id = met.unit_id " +
			" WHERE sens.sensor_id = " + strconv.Itoa(id) +
			" AND rtime LIKE '%" + date + "%' " + //TODO: sanitize date string
			" GROUP BY u.unit_id " +
			" ORDER BY mes.rvalue DESC")
		if err != nil {
			return nil, err

		}
		var (
			rvalueMin  sql.NullFloat64
			rvalueMax  sql.NullFloat64
			rtime	   sql.NullString
			unitName   sql.NullString
			metricName sql.NullString
			precision  sql.NullInt32
			dataPoints []DataPointMinMax
		)
		for rows.Next() {
			err := rows.Scan(&rvalueMax, &rvalueMin, &unitName, &metricName, &precision, &rtime)

			if err != nil {
				return nil, err
			}

			// only fetch if value is valid
			if rvalueMin.Valid {
				dataPoints = append(dataPoints, DataPointMinMax{
					Name:      metricName.String,
					Unit:      unitName.String,
					Precision: precision.Int32,
					Min:       rvalueMin.Float64,
					Time:      rtime.String[0:10], // trim timestamp, leave only date
					Max:       rvalueMax.Float64,
				})
			}
		}
		data = append(data, SensorMinMax{
			SensorId: id,
			Name:     name,
			SerialCode: serialCode,
			Data: dataPoints,
		})
	}
	return data, nil
}