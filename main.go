package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct {
	db *sql.DB
	router *mux.Router
}

func (s *server) Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/sensors", SensorIndex)
	router.HandleFunc("/sensors/{date}", SensorStats)
	s.router = router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(db *sql.DB) *server {
	s := &server{db: db}
	s.Routes()
	return s
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()
	server := newServer(db)

	fmt.Fprintf(os.Stdout, "%s\n", "Listening on port 8080... Ctrl+C to stop")
	return http.ListenAndServe(":8080", server.router)
}
