package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// katra sensora katras metrikas minimālā un maksimālā vērtība pieprasījuma parametrā norādītajā datumā.

func run() error {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/sensors", SensorIndex)
	router.HandleFunc("/sensors/{date}", SensorStats)

	fmt.Fprintf(os.Stdout, "%s\n", "Listening on port 8080... Ctrl+C to stop")
	return http.ListenAndServe(":8080", router)

	// DB stuff

	// db, err := sql.Open("sqlite3", "./db/aranet.db")
	// if err != nil {

	// 	return err
	// }
	// rows, err := db.Query("SELECT sensor_id, serial_code, `name` FROM sensors")
	// if err != nil {
	// 	return err

	// }
	// var (
	// 	sensorID   int
	// 	serialCode int
	// 	name       string
	// )

	// for rows.Next() {
	// 	err := rows.Scan(&sensorID, &serialCode, &name)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(strconv.Itoa(sensorID) + " (" + strconv.Itoa(serialCode) + "): " + name)
	// }
}

// Index - API home page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SQLite+Go practice example API v1")
}

// SensorIndex - fetch newest reading from every sensor
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sensor Index!")
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
	// fmt.Fprintln(w, "fetching for date: ", date)

}
