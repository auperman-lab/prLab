package cmd

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	//subrouter := router.PathPrefix("/api/v1").Subrouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	slog.Info("Listening on", "addr", s.addr)

	return http.ListenAndServe(s.addr, router)
}