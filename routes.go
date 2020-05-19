package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Routes - generate routes
func (s *server) Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", logRequest(Index))
	router.HandleFunc("/sensors", logRequest(jsonHeader(SensorIndex)))
	router.HandleFunc("/sensors/{date}", logRequest(jsonHeader(SensorStats)))
	s.router = router
}

// logRequest - middleware - logs any incoming requests to std.OUT
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s:%s  FROM:%s",r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

// jsonHeader - middleware - adds content type to response headers
func jsonHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

// Index - API home page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SQLite+Go practice example API v1")
}

// SensorIndex - fetch newest reading from every sensor
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	data, err := GetAllSensors(DB)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
}

// SensorStats - fetch min and max sensor values on specific date
func SensorStats(w http.ResponseWriter, r *http.Request) {
	// query := mux.Vars(r)
	// date := query["date"]
	//data := Sensors{
	//	Sensor{Name: "sensor1", Val1: 142, Val2: 152},
	//	Sensor{Name: "sensor2", Val1: 14, Val2: 15},
	//	Sensor{Name: "sensor3", Val1: 12, Val2: 12},
	//}
	// TODO: implement this endpoint
	json.NewEncoder(w).Encode("not yet implemented")
}