package product

import (
	"net/http"

	"github.com/Likhon22/ecom/rest/middleware"
)

func (h *Handler) ProductRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("GET /product/{id}", mngr.With(http.HandlerFunc(h.GetProductByID)))
	mux.Handle("POST /product", mngr.With(http.HandlerFunc(h.CreateProduct), middleware.Authenticate))
	mux.Handle("PATCH /product/{id}", mngr.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /product", mngr.With(http.HandlerFunc(h.DeleteProduct)))
}
