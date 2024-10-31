package http

import (
	ctrl "github.com/auperman-lab/lab2/internal/controller/http"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, productCtrl *ctrl.ProductController) {

	productsRouter := r.PathPrefix("/products").Subrouter()

	productsRouter.HandleFunc("", productCtrl.CreateProduct).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id:[0-9]+}", productCtrl.GetProductByID).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{name}", productCtrl.GetProductByName).Methods(http.MethodGet)
	productsRouter.HandleFunc("", productCtrl.UpdateProduct).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id:[0-9]+}", productCtrl.DeleteProductByID).Methods(http.MethodDelete)
	productsRouter.HandleFunc("/{page:[0-9]+}/{limit:[0-9]+}", productCtrl.GetAllProducts).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id:[0-9]+}/upload", productCtrl.UpdateProductImage).Methods(http.MethodPut)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("\nHello, This is an api that presents linella products !\n"))
		w.Write([]byte("\n" +
			" _    _          _ _\n" +
			"| |  (_)_ _  ___| | |__ _\n" +
			"| |__| | ' \\/ -_| | / _` |\n" +
			"|____|_|_||_\\___|_|_\\__,_|\n" +
			"\n"))

	})

}
