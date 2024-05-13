package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

type server struct {
	dbConn *sql.DB
}

func NewServer(dbConn *sql.DB) *server {
	return &server{
		dbConn: dbConn,
	}
}

func (s *server) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	_, err := s.dbConn.ExecContext(r.Context(), "INSERT INTO users (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Hello, %s!", name)
	w.Write([]byte(msg))
}
