package rest

import (
	"net/http"

	"github.com/Likhon22/ecom/rest/handlers"
	"github.com/Likhon22/ecom/rest/middleware"
)

func initRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /product/{id}", mngr.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("POST /product",mngr.With(http.HandlerFunc(handlers.CreateProduct), middleware.Authenticate))
	mux.Handle("PATCH /product/{id}", mngr.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /product", mngr.With(http.HandlerFunc(handlers.DeleteProduct)))
	mux.Handle("POST /users", mngr.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /login", mngr.With(http.HandlerFunc(handlers.Login)))
}
