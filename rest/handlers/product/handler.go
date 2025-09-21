package product

import (
	"github.com/Likhon22/ecom/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	service     ProductService
}

func NewHandler(middlewares *middleware.Middlewares, ProductService ProductService) *Handler {
	return &Handler{
		middlewares: middlewares,
		service:     ProductService,
	}
}
