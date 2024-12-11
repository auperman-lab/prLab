package http

import (
	ctrl "github.com/auperman-lab/lab2/internal/controller/http"
	"github.com/auperman-lab/lab2/internal/middleware"
	repo "github.com/auperman-lab/lab2/internal/repository"
	svc "github.com/auperman-lab/lab2/internal/service"
	"github.com/auperman-lab/lab2/raft"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type APIServer struct {
	addr string
	db   *gorm.DB
	node *raft.Node
}

func NewAPIServer(addr string, db *gorm.DB, node *raft.Node) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		node: node,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.Use(middleware.LeaderCheckerMiddleware(s.node.GetLeader))
	router.Use(middleware.RaftReplicationMiddleware(s.node.AppendLogs))

	productRepo := repo.NewProductRepository(s.db)
	productSvc := svc.NewProductService(productRepo)
	productCtrl := ctrl.NewProductController(productSvc)
	RegisterRoutes(router, productCtrl)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	slog.Info("Listening on", "addr", s.addr)

	return http.ListenAndServe(s.addr, router)
}
