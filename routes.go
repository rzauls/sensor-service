package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Index - API home page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SQLite+Go practice example API v1")
}

// SensorIndex - fetch newest reading from every sensor
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	//GetAllSensors(db)
	json.NewEncoder(w).Encode("test")
}

// SensorStats - fetch min and max sensor values on specific date
func SensorStats(w http.ResponseWriter, r *http.Request) {
	// query := mux.Vars(r)
	// date := query["date"]
	data := Sensors{
		Sensor{Name: "sensor1", Val1: 142, Val2: 152},
		Sensor{Name: "sensor2", Val1: 14, Val2: 15},
		Sensor{Name: "sensor3", Val1: 12, Val2: 12},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}