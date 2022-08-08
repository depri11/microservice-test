package auth

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func NewRoute(r *mux.Router, db *sql.DB) {
	route := r.PathPrefix("/auth").Subrouter()

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	route.HandleFunc("/", handler.Register).Methods("POST")

}
