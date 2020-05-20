package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	db *sql.DB
	router *mux.Router
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
		log.Fatalf("%s\n", err)
	}
}

func run() error {
	if err := InitDB(); err != nil {
		return err
	}
	defer DB.Close()

	server := newServer(DB)
	log.Printf("Listening on port 8080...")
	return http.ListenAndServe(":8080", server.router)
}
