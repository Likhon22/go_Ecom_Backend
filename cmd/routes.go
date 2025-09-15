package cmd

import (
	"net/http"

	"github.com/Likhon22/ecom/handlers"
	"github.com/Likhon22/ecom/middleware"
)

func initRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /product/{id}", mngr.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("POST /product", mngr.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("PATCH /product", mngr.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /product", mngr.With(http.HandlerFunc(handlers.DeleteProduct)))
}
