package http

import (
	ctrl "github.com/auperman-lab/lab2/internal/controller/http"
	repo "github.com/auperman-lab/lab2/internal/repository"
	svc "github.com/auperman-lab/lab2/internal/service"
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
	productRepo := repo.NewProductRepository(s.db)
	productSvc := svc.NewProductService(productRepo)
	productCtrl := ctrl.NewProductController(productSvc)
	RegisterRoutes(router, productCtrl)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	slog.Info("Listening on", "addr", s.addr)

	return http.ListenAndServe(s.addr, router)
}
