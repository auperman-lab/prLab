package cmd

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type APIServer struct {
	addr string
	//db   *sql.DB
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
		//db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	//subrouter := router.PathPrefix("/api/v1").Subrouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	slog.Info("Listening on", "addr", s.addr)

	return http.ListenAndServe(s.addr, router)
}
