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
	if err := InitDB(); err != nil {
		return err
	}
	defer DB.Close()

	server := newServer(DB)
	fmt.Fprintf(os.Stdout, "%s\n", "Listening on port 8080... Ctrl+C to stop")
	return http.ListenAndServe(":8080", server.router)
}
